// Code generated by zenrpc; DO NOT EDIT.

package testdata

import (
	"context"
	"encoding/json"

	"github.com/semrush/zenrpc"
	"github.com/semrush/zenrpc/smd"
)

var RPC = struct {
	ArithService struct{ Sum, Positive, DoSomething, Multiply, Divide, Pow, Pi, SumArray string }
	PhoneBook    struct{ Get, ValidateSearch, ById, Delete, Remove, Save string }
}{
	ArithService: struct{ Sum, Positive, DoSomething, Multiply, Divide, Pow, Pi, SumArray string }{
		Sum:         "sum",
		Positive:    "positive",
		DoSomething: "dosomething",
		Multiply:    "multiply",
		Divide:      "divide",
		Pow:         "pow",
		Pi:          "pi",
		SumArray:    "sumarray",
	},
	PhoneBook: struct{ Get, ValidateSearch, ById, Delete, Remove, Save string }{
		Get:            "get",
		ValidateSearch: "validatesearch",
		ById:           "byid",
		Delete:         "delete",
		Remove:         "remove",
		Save:           "save",
	},
}

func (ArithService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Sum": {
				Description: `Sum sums two digits and returns error with error code as result and IP from context.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Boolean,
				},
			},
			"Positive": {
				Description: ``,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Boolean,
				},
			},
			"DoSomething": {
				Description: ``,
				Parameters:  []smd.JSONSchema{},
			},
			"Multiply": {
				Description: `Multiply multiples two digits and returns result.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Integer,
				},
			},
			"Divide": {
				Description: `Divide divides two numbers.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
				},
			},
			"Pow": {
				Description: `Pow returns x**y, the base-x exponential of y. If Exp is not set then default value is 2.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "base",
						Optional:    false,
						Description: ``,
						Type:        smd.Float,
					},
					{
						Name:        "exp",
						Optional:    true,
						Description: ``,
						Type:        smd.Float,
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Float,
				},
			},
			"Pi": {
				Description: `PI returns math.Pi.`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Float,
				},
			},
			"SumArray": {
				Description: `SumArray returns sum all items from array`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "array",
						Optional:    true,
						Description: ``,
						Type:        smd.Array,
						Items: map[string]string{
							"type": smd.Float,
						},
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Float,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s ArithService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.ArithService.Sum:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Sum(ctx, args.A, args.B))

	case RPC.ArithService.Positive:
		resp.Set(s.Positive())

	case RPC.ArithService.DoSomething:
		s.DoSomething()

	case RPC.ArithService.Multiply:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Multiply(args.A, args.B))

	case RPC.ArithService.Divide:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Divide(args.A, args.B))

	case RPC.ArithService.Pow:
		var args = struct {
			Base float64  `json:"base"`
			Exp  *float64 `json:"exp"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"base", "exp"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		//zenrpc:exp:2 	exponent could be empty
		if args.Exp == nil {
			var v float64 = 2
			args.Exp = &v
		}

		resp.Set(s.Pow(args.Base, args.Exp))

	case RPC.ArithService.Pi:
		resp.Set(s.Pi())

	case RPC.ArithService.SumArray:
		var args = struct {
			Array *[]float64 `json:"array"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"array"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		//zenrpc:array:[]float64{1,2,4}
		if args.Array == nil {
			var v []float64 = []float64{1, 2, 4}
			args.Array = &v
		}

		resp.Set(s.SumArray(args.Array))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (PhoneBook) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Get": {
				Description: `Get returns all people from DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "search",
						Optional:    false,
						Description: ``,
						Type:        smd.Object,
					},
					{
						Name:        "page",
						Optional:    true,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "count",
						Optional:    true,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Array,
					Items: map[string]string{
						"type": smd.Object,
					},
				},
			},
			"ValidateSearch": {
				Description: `ValidateSearch returns given search as result.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "search",
						Optional:    true,
						Description: ``,
						Type:        smd.Object,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
				},
			},
			"ById": {
				Description: `ById returns Person from DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
				},
			},
			"Delete": {
				Description: `Delete marks person as deleted.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Boolean,
				},
			},
			"Remove": {
				Description: `Removes deletes person from DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{Description: `operation result`,
					Optional: false,
					Type:     smd.Boolean,
				},
			},
			"Save": {
				Description: `Save saves person to DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "p",
						Optional:    false,
						Description: ``,
						Type:        smd.Object,
					},
					{
						Name:        "replace",
						Optional:    true,
						Description: ``,
						Type:        smd.Boolean,
					},
				},
				Returns: smd.JSONSchema{
					Optional: false,
					Type:     smd.Integer,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s PhoneBook) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.PhoneBook.Get:
		var args = struct {
			Search PersonSearch `json:"search"`
			Page   *int         `json:"page"`
			Count  *int         `json:"count"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"search", "page", "count"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		//zenrpc:count:50 page size
		if args.Count == nil {
			var v int = 50
			args.Count = &v
		}

		//zenrpc:page:0 current page
		if args.Page == nil {
			var v int = 0
			args.Page = &v
		}

		resp.Set(s.Get(args.Search, args.Page, args.Count))

	case RPC.PhoneBook.ValidateSearch:
		var args = struct {
			Search *PersonSearch `json:"search"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"search"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.ValidateSearch(args.Search))

	case RPC.PhoneBook.ById:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.ById(args.Id))

	case RPC.PhoneBook.Delete:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Delete(args.Id))

	case RPC.PhoneBook.Remove:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Remove(args.Id))

	case RPC.PhoneBook.Save:
		var args = struct {
			P       Person `json:"p"`
			Replace *bool  `json:"replace"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"p", "replace"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		//zenrpc:replace:false update person if exist
		if args.Replace == nil {
			var v bool = false
			args.Replace = &v
		}

		resp.Set(s.Save(args.P, args.Replace))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
