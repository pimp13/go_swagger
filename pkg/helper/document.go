package helper

/*
type EndpointDoc struct {
	Method      string
	Path        string
	Summary     string
	Description string
}

func Document(method, path, summary, description string, v ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiDocs := make(map[string]any)
		key := method + ":" + path
		apiDocs[key] = &EndpointDoc{
			Method:      method,
			Path:        path,
			Summary:     summary,
			Description: description,
		}
		c.Next()
	}
}
*/

/*
type SwaggerHelper struct {
	docs     *swag.Spec
	apiInfo  map[string]interface{}
	paths    map[string]interface{}
	defs     map[string]interface{}
	basePath string
}

func NewSwaggerHelper(title, description, version string) *SwaggerHelper {
	return &SwaggerHelper{
		docs: &swag.Spec{
			Version:          "",
			Host:             "",
			BasePath:         "",
			Schemes:          []string{},
			Title:            "",
			Description:      "",
			InfoInstanceName: "",
			SwaggerTemplate:  "",
		},
		apiInfo: map[string]interface{}{
			"title":       title,
			"description": description,
			"version":     version,
		},
		paths:    make(map[string]interface{}),
		defs:     make(map[string]interface{}),
		basePath: "/api/v1",
	}
}

// AutoGenerateRouteDocs مستندات را به صورت خودکار تولید می‌کند
func (sh *SwaggerHelper) AutoGenerateRouteDocs(router *gin.Engine) {
	for _, route := range router.Routes() {
		pathKey := sh.normalizePath(route.Path)
		pathItem := map[string]interface{}{
			strings.ToLower(route.Method): sh.generateOperation(route),
		}

		if existing, exists := sh.paths[pathKey]; exists {
			if pathMap, ok := existing.(map[string]interface{}); ok {
				pathMap[strings.ToLower(route.Method)] = pathItem[strings.ToLower(route.Method)]
				sh.paths[pathKey] = pathMap
			}
		} else {
			sh.paths[pathKey] = pathItem
		}
	}
}

// normalizePath مسیرها را استاندارد می‌کند
func (sh *SwaggerHelper) normalizePath(path string) string {
	if !strings.HasPrefix(path, sh.basePath) {
		path = sh.basePath + path
	}
	return path
}

// generateOperation اطلاعات عملیات را تولید می‌کند
func (sh *SwaggerHelper) generateOperation(route gin.RouteInfo) map[string]interface{} {
	handlerName := runtime.FuncForPC(reflect.ValueOf(route.HandlerFunc)).Name()
	summary := fmt.Sprintf("Auto-generated endpoint for %s", route.Path)

	return map[string]interface{}{
		"summary":     summary,
		"description": fmt.Sprintf("Handler: %s", handlerName),
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Successful operation",
			},
		},
	}
}

// RegisterResponseModel مدل پاسخ را ثبت می‌کند
func (sh *SwaggerHelper) RegisterResponseModel(name string, model interface{}) {
	schema := sh.generateSchema(model)
	sh.defs[name] = schema
}

// generateSchema ساختار Swagger Schema را تولید می‌کند
func (sh *SwaggerHelper) generateSchema(model interface{}) map[string]interface{} {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := map[string]interface{}{
		"type": "object",
	}

	properties := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		fieldName := strings.Split(jsonTag, ",")[0]
		fieldType := sh.mapTypeToSwagger(field.Type)

		properties[fieldName] = map[string]interface{}{
			"type": fieldType,
		}
	}

	schema["properties"] = properties
	return schema
}

// mapTypeToSwagger نوع Go را به نوع Swagger تبدیل می‌کند
func (sh *SwaggerHelper) mapTypeToSwagger(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.Struct:
		return "object"
	case reflect.Slice, reflect.Array:
		return "array"
	default:
		return "string"
	}
}

// GenerateSpec مستندات نهایی را تولید می‌کند
func (sh *SwaggerHelper) GenerateSpec() string {
	swagger := map[string]interface{}{
		"swagger": "2.0",
		"info": map[string]interface{}{
			"title":       sh.apiInfo["title"],
			"description": sh.apiInfo["description"],
			"version":     sh.apiInfo["version"],
		},
		"basePath":    sh.basePath,
		"paths":       sh.paths,
		"definitions": sh.defs,
	}

	jsonData, _ := json.MarshalIndent(swagger, "", "  ")
	return string(jsonData)
}
*/
