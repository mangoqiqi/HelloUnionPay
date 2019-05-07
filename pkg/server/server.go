package server

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/unionpay/common"
	"github.com/unionpay/config"
	"github.com/unionpay/pkg/sql"
	"github.com/vinkdong/gox/log"
	"net/http"
	"reflect"
)

type Server struct {
	Addr      string
	ServerMux *http.ServeMux
	Server    *http.Server
}

var withhold common.PayForAnother
var engine *xorm.Engine

func NewServer(port int) *Server {
	s := &Server{
		Addr:      fmt.Sprintf(":%d", port),
		ServerMux: http.NewServeMux(),
	}
	return s
}

func (s *Server) HandlerInit() {
	log.Info("Api /v1.0/Withholding init ..")
	s.ServerMux.HandleFunc("/v1.0/Withholding", withHolding) //单笔代付
	log.Info("Api /v1.0/MuxWithholding init ..")
	s.ServerMux.HandleFunc("/v1.0/MuxWithholding", muxWithHolding) //多笔代付
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:    s.Addr,
		Handler: s.ServerMux,
	}
	s.Server = server
	log.Infof("server starting on %s", s.Addr)
	return server.ListenAndServe()
}

func StartServers(port int) error {
	//init config
	log.Info("init config ...")
	err := config.InitConfig()
	if err != nil {
		return err
	}
	withhold, err = common.NewPayForAnother(common.ApiConfig{
		Url: "https://gateway.test.95516.com",
	})
	if err != nil {
		return err
	}
	//init mysql
	log.Info("init mysql ...")
	engine, err = sql.NewEngine("test", "root","root","192.168.56.11","3306")
	if err != nil {
		return err
	}
	err = engine.Ping()
	if err !=nil {
		return err
	}
	s := NewServer(port)
	s.HandlerInit()
	return s.Start()
}

func withHolding(w http.ResponseWriter, r *http.Request) {

	//TODO: json formatch
	customer := common.CustomerInfo{
		CustomerNm: "全渠道",
		PhoneNo:    "13552535506",
	}
	read := common.RequestParams{
		AccNo:    "6216261000000000018",
		Customer: &customer,
	}
	result, err := withhold.Pay(11115, &read)
	if err != nil {
		log.Error(err)
		return
	}
	//TODO:real name check
	//TODO:storage result
	v := reflect.ValueOf(result)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			fmt.Println(key.Interface(), val.Interface())
		}
	}
	w.Write([]byte(`{"success":true}`))
}

func muxWithHolding(w http.ResponseWriter, r *http.Request)  {
	//TODO:multi Withholdings
}
