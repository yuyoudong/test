package demo

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/errorcode"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/form_validator"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/log"
	domain "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/domain/demo"
	"github.com/gin-gonic/gin"
	"github.com/jinguoxing/af-go-frame/core/transport/rest/ginx"
)

type Service struct {
	uc domain.UseCase
}

func NewService(uc domain.UseCase) *Service {
	return &Service{uc: uc}
}

func (s *Service) Create(c *gin.Context) {
	req := &domain.CreateReqParam{}
	log.Info("requst sql masking")
	err2 := c.ShouldBindJSON(&req.CreateReqBodyParam)
	if err2 != nil {
	}
	data := c.GetHeader("Content-type")
	if strings.ToLower(data) != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Content-type参数校验不通过,必须为json"})
		return
	}

	_, err := form_validator.BindJsonAndValid(c, &req.CreateReqBodyParam)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		if errors.As(err, &form_validator.ValidErrors{}) {
			ginx.ResErrJson(c, errorcode.Detail(errorcode.PublicInvalidParameter, err))
			return
		}
	}

	output := bytes.Buffer{}
	output.WriteString("SELECT ")
	for i := 0; i < len(req.CreateReqBodyParam.Fields); i++ {

		field := req.CreateReqBodyParam.Fields[i].Field //chinese_name string, sensitive int, classified
		chinese_name := req.CreateReqBodyParam.Fields[i].ChineseName
		sensitive := req.CreateReqBodyParam.Fields[i].Sensitive
		classified := req.CreateReqBodyParam.Fields[i].Classified
		field_type := req.CreateReqBodyParam.Fields[i].FieldType
		if !check_string_input(field, c, "field") {
			return
		}
		if !check_string_input(chinese_name, c, "chinese_name") {
			return
		}
		if !check_string_input(field_type, c, "field_type") {
			return
		}

		if !check_int_input(classified, c, "classified") {
			return
		}
		if !check_int_input(sensitive, c, "sensitive") {
			return
		}

		field_str := str_concat(field_type, field, chinese_name, sensitive, classified)
		output.WriteString(field_str)

		if i < len(req.CreateReqBodyParam.Fields)-1 {
			output.WriteString(",")
		}
	}
	if !check_string_input(req.CreateReqBodyParam.TableName, c, "table_name") {
		return
	}
	output.WriteString(" ")
	output.WriteString("FROM ")
	output.WriteString(req.CreateReqBodyParam.TableName)
	// output.WriteString(";")

	sql_masking := map[string]string{ //初始化返回值，类型为map
		"masked_sql": string(output.String()),
	}
	log.Info("new sql:")
	log.Info(output.String())
	ginx.ResOKJson(c, sql_masking)

}
func Containss(slice []any, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}
func Contains(slice []int, element int) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}
func check_int_input(check_int int, c *gin.Context, field string) (result bool) {
	code := "DataMasking.Form.InvalidParameter"
	solution := "请使用请求参数构造规范化的请求字符串,详细信息参见产品API文档"
	err_json := "参数值校验不通过:json格式错误"
	detail_array := []map[string]string{}
	int_c := []int{0, 1}
	if Contains(int_c, check_int) && fmt.Sprintf("%T", check_int) == "int" {
		return true
	} else {
		detail_array = append(detail_array, map[string]string{"key": field, "meggage": fmt.Sprintf("%s为必填字段输入类型必须为int且只能为0、1", field)})
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "description": err_json, "solution": solution, "message": detail_array})
		return false
	}

}

func check_string_input(check_string string, c *gin.Context, field string) (result bool) {
	code := "DataMasking.Form.InvalidParameter"
	solution := "请使用请求参数构造规范化的请求字符串,详细信息参见产品API文档"
	err_json := "参数值校验不通过:json格式错误"
	err_input := "参数值校验不通过"
	compile_zh := regexp.MustCompile("[\u4e00-\u9fa5]")
	compile1 := regexp.MustCompile("\\s")
	spcChar := []string{`,`, `?`, `*`, `|`, `{`, `}`, `\`, `/`, `$`, `、`, `·`, "`", `'`, `"`, `#`, `!`, `^`}
	detail_array := []map[string]string{}

	if check_string == " " || check_string == "" {
		detail_array = append(detail_array, map[string]string{"key": field, "meggage": fmt.Sprintf("%s为必填字段输入类型必须为string且不能为空或空格", field)})
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "description": err_json, "solution": solution, "message": detail_array})
		return false
	}

	if len(check_string) == 0 {
		detail_array = append(detail_array, map[string]string{"key": field, "meggage": fmt.Sprintf("%s为必填字段输入类型必须为string且不能为空", field)})
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "description": err_json, "solution": solution, "message": detail_array})
		return false

	} else if compile_zh.MatchString(check_string) && field != "chinese_name" {

		detail_array = append(detail_array, map[string]string{"key": field, "meggage": fmt.Sprintf("%s为包含中文", field)})
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "description": err_input, "solution": solution, "message": detail_array})
		return false
	} else if compile1.MatchString(check_string) && field != "chinese_name" && field != "field" {
		detail_array = append(detail_array, map[string]string{"key": field, "meggage": fmt.Sprintf("%s不能包含空格", field)})
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "description": err_input, "solution": solution, "message": detail_array})
		return false
	} else if strings.ContainsAny(check_string, strings.Join(spcChar, "")) && field != "chinese_name" {
		detail_array = append(detail_array, map[string]string{"key": field, "meggage": fmt.Sprintf("%s不能包含特殊字符", field)})
		c.JSON(http.StatusBadRequest, gin.H{"code": code, "description": err_input, "solution": solution, "message": detail_array})
		return false
	} else {
		return true
	}
}

var user_defined_masking_rule = map[string]string{
	"姓名":   "MASK_LAST_1",
	"电话号码": "MASK_MID_5,8",
	"身份证":  "MASK_MID_7,14",
	"护照":   "MASK_LAST_4",
}

func str_concat(field_type string, field string, chinese_name string, sensitive int, classified int) string {
	if strings.ToLower(field_type) != "string" {
		return fmt.Sprintf(`"` + field + `"`)
	}
	_, ok := user_defined_masking_rule[chinese_name]
	if ok == true {
		if chinese_name == "姓名" { //CONCAT(SUBSTR(%s,1,1),'*')
			return fmt.Sprintf("CONCAT(SUBSTR(" + `"` + field + `"` + ",1,1),'**')  AS " + `"` + field + `"`)
			// return fmt.Sprintf("CONCAT(SUBSTR(%s,1,1),'**')  AS %s", field, field)
		} else if chinese_name == "电话号码" { //中间4位用*代替
			return fmt.Sprintf("CONCAT(SUBSTR(" + `"` + field + `"` + ",1,3),'****',SUBSTR(" + `"` + field + `"` + ",8,12)) AS " + `"` + field + `"`)
			// return fmt.Sprintf("CONCAT(SUBSTR(%s,1,3),'****',SUBSTR(%s,8,12)) AS %s", field, field, field)
		} else if chinese_name == "身份证" { // 中间的生日8位用*代替
			return fmt.Sprintf("CONCAT(SUBSTR(" + `"` + field + `"` + ",1,6),'********',SUBSTR(" + `"` + field + `"` + ",15,18)) AS " + `"` + field + `"`)
			//return fmt.Sprintf("CONCAT(SUBSTR(%s,1,6),'********',SUBSTR(%s,15,18)) AS %s", field, field, field)
		} else if chinese_name == "护照" {
			return fmt.Sprintf("CONCAT(SUBSTR(" + `"` + field + `"` + ",1,5),'****')  AS " + `"` + field + `"`)
		} else {
			return ""
		}
	} else {
		base_result := sensitive + classified
		if base_result == 0 { //不敏感、不涉密，不处理
			return field
		} else if base_result == 1 { //敏感、涉密，字段全部进行脱敏
			return fmt.Sprintf("rpad('',6,'*') as " + `"` + field + `"`)
			// return fmt.Sprintf("CASE WHEN LENGTH(%s) > 2 THEN  RPAD(SUBSTR(%s,1,LENGTH(%s)/2),LENGTH(%s),'*') ELSE '*' END AS %s", field, field, field, field, field)
		} else { //敏感、涉密，字段全部进行脱敏
			return fmt.Sprintf("rpad('',6,'*') as " + `"` + field + `"`)
		}
	}
	return ""
}
