package openpid

import (
	"github.com/openimchat/weapp/v3/request"
)

const apiGetOpenpId = "/wxa/getpluginopenpid"

type GetOpenpIdRequest struct {
	Code string `json:"code"`
}

type GetOpenpIdResponse struct {
	request.CommonError
	OpenpId string `json:"openpid"`
}

// code换取用户手机号。 每个code只能使用一次，code的有效期为5min
func (cli *OpenPid) GetOpenpId(req *GetOpenpIdRequest) (*GetOpenpIdResponse, error) {

	api, err := cli.combineURI(apiGetOpenpId, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetOpenpIdResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
