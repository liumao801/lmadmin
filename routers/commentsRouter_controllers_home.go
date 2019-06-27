package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["liumao801/lmadmin/controllers/home:ErrorsController"] = append(beego.GlobalControllerRouter["liumao801/lmadmin/controllers/home:ErrorsController"],
        beego.ControllerComments{
            Method: "Page404",
            Router: `/404`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liumao801/lmadmin/controllers/home:ErrorsController"] = append(beego.GlobalControllerRouter["liumao801/lmadmin/controllers/home:ErrorsController"],
        beego.ControllerComments{
            Method: "Page4042",
            Router: `/404-2`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liumao801/lmadmin/controllers/home:ErrorsController"] = append(beego.GlobalControllerRouter["liumao801/lmadmin/controllers/home:ErrorsController"],
        beego.ControllerComments{
            Method: "Page500",
            Router: `/500`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
