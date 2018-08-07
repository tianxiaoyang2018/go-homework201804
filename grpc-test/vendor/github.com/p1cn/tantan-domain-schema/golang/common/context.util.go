package common

func UpdateServiceContext(serviceContext *Context, serviceName string) *Context {
	if serviceContext == nil {
		serviceContext = &Context{}
	}
	serviceContext.ServiceTrace = append(serviceContext.ServiceTrace, serviceContext.Service)
	serviceContext.Service = serviceName
	return serviceContext
}

