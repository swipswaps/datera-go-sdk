package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type BootDrive struct {
	Path      string   `json:"path,omitempty" mapstructure:"path"`
	Causes    []string `json:"causes,omitempty" mapstructure:"causes"`
	Health    string   `json:"health,omitempty" mapstructure:"health"`
	Id        string   `json:"id,omitempty" mapstructure:"id"`
	OpState   string   `json:"op_state,omitempty" mapstructure:"op_state"`
	Size      int      `json:"size,omitempty" mapstructure:"size"`
	SlotLabel string   `json:"slot_label,omitempty" mapstructure:"slot_label"`
}

type BootDrives struct {
	Path string
}

func newBootDrives(path string) *BootDrives {
	return &BootDrives{
		Path: _path.Join(path, "boot_drives"),
	}
}

type BootDrivesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *BootDrives) List(ro *BootDrivesListRequest) ([]*BootDrive, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap()}
	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*BootDrive{}
	for _, data := range rs.Data {
		elem := &BootDrive{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type BootDrivesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"-"`
}

func (e *BootDrives) Get(ro *BootDrivesGetRequest) (*BootDrive, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &BootDrive{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
