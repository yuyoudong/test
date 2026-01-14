package demo_test

import (
	"bytes"
	"context"
	stdJson "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/adapter/controller"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/adapter/controller/demo/v1"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/errorcode"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/form_validator"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/log"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/util"
	domain "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/domain/demo"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinguoxing/af-go-frame/core/transport/rest"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	preUrl = "/api/data-masking/v1"
)

var (
	uc     = &demoUseCase{}
	engine *gin.Engine
)

type demoUseCase struct {
}

func (f *demoUseCase) Create(ctx context.Context, req *domain.CreateReqParam) (*domain.CreateRespParam, error) {
	// TODO implement me
	panic("implement me")
}

func TestMain(m *testing.M) {
	log.InitProjectLogger()
	// 初始化验证器
	err := form_validator.SetupValidator()
	if err != nil {
		panic(err)
	}

	// patches := gomonkey.ApplyFuncReturn(log.Infof)
	// defer patches.Reset()
	// patches.ApplyFuncReturn(log.Info)
	// patches.ApplyFuncReturn(log.Error)
	// patches.ApplyFuncReturn(log.Errorf)
	//
	// zapxLog := &zapx.ZapWriter{}
	// patches.ApplyMethodReturn(zapxLog, "Info")

	setupRouter()

	os.Exit(m.Run())
}

func setupRouter() {
	r := &controller.Router{
		DemoApi: demo.NewService(uc),
	}

	engine = controller.NewHttpEngine(r)
}

type utData struct {
	uriParam      []any
	queryParam    map[string]string
	formBodyParam url.Values
	jsonBodyParam any
	statusCode    int
	errCode       string

	needReq []any
}

func controllerUTFunc(method, uriPattern string, d *utData) {
	uri := uriPattern
	if len(d.uriParam) > 0 {
		uri = fmt.Sprintf(uri, d.uriParam...)
	}

	params := make([]string, 0, len(d.queryParam))
	for k, v := range d.queryParam {
		params = append(params, k+"="+v)
	}
	if len(params) > 0 {
		uri = uri + "?" + strings.Join(params, "&")
	}

	if d.formBodyParam != nil && d.jsonBodyParam != nil {
		panic("only form body or json body")
	}

	var body io.Reader
	if d.jsonBodyParam != nil {
		if jsonStr, ok := d.jsonBodyParam.(string); ok {
			body = strings.NewReader(jsonStr)
		} else {
			b, err := stdJson.Marshal(d.jsonBodyParam)
			So(err, ShouldBeNil)
			body = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequest(method, uri, body)
	So(err, ShouldBeNil)

	rec := httptest.NewRecorder()
	engine.ServeHTTP(rec, req)

	resp := rec.Result()
	So(resp.StatusCode, ShouldEqual, d.statusCode)

	if d.statusCode == http.StatusOK {
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	So(err, ShouldBeNil)
	defer resp.Body.Close()

	errResp := &rest.HttpError{}
	err = stdJson.Unmarshal(respBody, errResp)
	So(err, ShouldBeNil)

	So(errResp.Code, ShouldEqual, d.errCode)
}

func TestService_Create(t *testing.T) {
	uri := preUrl + "/demos/%v"
	method := http.MethodPost

	Convey("TestService_Create", t, func() {
		Convey("RequestParamUT", func() {
			patches := gomonkey.ApplyMethodReturn(uc, "Create", &domain.CreateRespParam{ID: "1", Name: "name"}, nil)
			defer patches.Reset()

			tests := map[string]*utData{
				"name has @": {
					uriParam: []any{"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"},
					queryParam: map[string]string{
						"p1": "1",
						"t2": "e1",
					},
					jsonBodyParam: map[string]string{
						"name": "n1@",
					},
					statusCode: http.StatusBadRequest,
					errCode:    errorcode.PublicInvalidParameter,
				},
			}

			for name, tt := range tests {
				Convey(name, func() {
					controllerUTFunc(method, uri, tt)
				})
			}
		})

		Convey("RequestHandleSucUT", func() {
			tests := map[string]*utData{
				"suc": {
					uriParam: []any{"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"},
					queryParam: map[string]string{
						"p1": "1",
						"t2": "e1",
					},
					jsonBodyParam: map[string]string{
						"name": "n1",
					},
					statusCode: http.StatusOK,
					needReq: []any{
						&domain.CreateReqParam{
							CreateReqPathParam: domain.CreateReqPathParam{
								PId: "4a5a3cc0-0169-4d62-9442-62214d8fcd8d",
							},
							CreateReqQueryParam: domain.CreateReqQueryParam{
								P1: util.ValueToPtr(1),
								P2: util.ValueToPtr("e1"),
							},
							CreateReqBodyParam: domain.CreateReqBodyParam{
								Name: util.ValueToPtr("n1"),
							},
						},
					},
				},
			}

			for name, tt := range tests {
				Convey(name, func() {
					patches := gomonkey.ApplyMethodFunc(uc, "Create", func(_ context.Context, req0 *domain.CreateReqParam) (*domain.CreateRespParam, error) {
						So(req0, ShouldResemble, tt.needReq[0])
						return nil, nil
					})
					defer patches.Reset()

					controllerUTFunc(method, uri, tt)
				})
			}
		})
	})
}
