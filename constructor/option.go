package constructor

import "errors"

// ResourcePoolConfig resource pool
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// ResourcePoolConfigOption option
type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

func SetMaxTotal(option *ResourcePoolConfigOption){

}

func NewResourcePoolConfig(name string,opts ...ResourcePoolConfigOptFunc)(*ResourcePoolConfig,error){
	if name==""{
		return nil,errors.New("name cannot be empty")
	}
	option := &ResourcePoolConfigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}
	for _, opt := range opts {
		opt(option)
	}
	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}

