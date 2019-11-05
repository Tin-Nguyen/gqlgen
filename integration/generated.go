package integration

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	models "github.com/99designs/gqlgen/integration/models-go"
	"github.com/99designs/gqlgen/integration/remote_api"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// region    ************************** generated!.gotpl **************************

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Element() ElementResolver
	Query() QueryResolver
	User() UserResolver
}

type DirectiveRoot struct {
	Magic func(ctx context.Context, obj interface{}, next graphql.Resolver, kind *int) (res interface{}, err error)
}

type ComplexityRoot struct {
	Element struct {
		Child      func(childComplexity int) int
		Error      func(childComplexity int) int
		Mismatched func(childComplexity int) int
	}

	Query struct {
		Complexity   func(childComplexity int, value int) int
		Date         func(childComplexity int, filter models.DateFilter) int
		Error        func(childComplexity int, typeArg *models.ErrorType) int
		JSONEncoding func(childComplexity int) int
		Path         func(childComplexity int) int
		Viewer       func(childComplexity int) int
	}

	User struct {
		Likes func(childComplexity int) int
		Name  func(childComplexity int) int
	}

	Viewer struct {
		User func(childComplexity int) int
	}
}

type ElementResolver interface {
	Child(ctx context.Context, obj *models.Element) (*models.Element, error)
	Error(ctx context.Context, obj *models.Element) (bool, error)
	Mismatched(ctx context.Context, obj *models.Element) ([]bool, error)
}
type QueryResolver interface {
	Path(ctx context.Context) ([]*models.Element, error)
	Date(ctx context.Context, filter models.DateFilter) (bool, error)
	Viewer(ctx context.Context) (*models.Viewer, error)
	JSONEncoding(ctx context.Context) (string, error)
	Error(ctx context.Context, typeArg *models.ErrorType) (bool, error)
	Complexity(ctx context.Context, value int) (bool, error)
}
type UserResolver interface {
	Likes(ctx context.Context, obj *remote_api.User) ([]string, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Element.child":
		if e.complexity.Element.Child == nil {
			break
		}

		return e.complexity.Element.Child(childComplexity), true

	case "Element.error":
		if e.complexity.Element.Error == nil {
			break
		}

		return e.complexity.Element.Error(childComplexity), true

	case "Element.mismatched":
		if e.complexity.Element.Mismatched == nil {
			break
		}

		return e.complexity.Element.Mismatched(childComplexity), true

	case "Query.complexity":
		if e.complexity.Query.Complexity == nil {
			break
		}

		args, err := ec.field_Query_complexity_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Complexity(childComplexity, args["value"].(int)), true

	case "Query.date":
		if e.complexity.Query.Date == nil {
			break
		}

		args, err := ec.field_Query_date_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Date(childComplexity, args["filter"].(models.DateFilter)), true

	case "Query.error":
		if e.complexity.Query.Error == nil {
			break
		}

		args, err := ec.field_Query_error_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Error(childComplexity, args["type"].(*models.ErrorType)), true

	case "Query.jsonEncoding":
		if e.complexity.Query.JSONEncoding == nil {
			break
		}

		return e.complexity.Query.JSONEncoding(childComplexity), true

	case "Query.path":
		if e.complexity.Query.Path == nil {
			break
		}

		return e.complexity.Query.Path(childComplexity), true

	case "Query.viewer":
		if e.complexity.Query.Viewer == nil {
			break
		}

		return e.complexity.Query.Viewer(childComplexity), true

	case "User.likes":
		if e.complexity.User.Likes == nil {
			break
		}

		return e.complexity.User.Likes(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "Viewer.user":
		if e.complexity.Viewer.User == nil {
			break
		}

		return e.complexity.Viewer.User(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `"This directive does magical things"
directive @magic(kind: Int) on FIELD_DEFINITION

type Element {
    child: Element!
    error: Boolean!
    mismatched: [Boolean!]
}

enum DATE_FILTER_OP {
    # multi
    # line
    # comment
    EQ
    NEQ
    GT
    GTE
    LT
    LTE
}

input DateFilter {
    value: String!
    timezone: String = "UTC"
    op: DATE_FILTER_OP = EQ
}

type Viewer {
    user: User
}

type Query {
    path: [Element]
    date(filter: DateFilter!): Boolean!
    viewer: Viewer
    jsonEncoding: String!
    error(type: ErrorType = NORMAL): Boolean!
    complexity(value: Int!): Boolean!
}

enum ErrorType {
    CUSTOM
    NORMAL
}

# this is a comment with a ` + "`" + `backtick` + "`" + `
`},
	&ast.Source{Name: "user.graphql", Input: `type User {
    name: String!
    likes: [String!]!
}
`},
)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) dir_magic_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["kind"]; ok {
		arg0, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["kind"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_complexity_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["value"]; ok {
		arg0, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["value"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_date_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 models.DateFilter
	if tmp, ok := rawArgs["filter"]; ok {
		arg0, err = ec.unmarshalNDateFilter2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilter(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["filter"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_error_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *models.ErrorType
	if tmp, ok := rawArgs["type"]; ok {
		arg0, err = ec.unmarshalOErrorType2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐErrorType(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["type"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Element_child(ctx context.Context, field graphql.CollectedField, obj *models.Element) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Element",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Element().Child(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*models.Element)
	fc.Result = res
	return ec.marshalNElement2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx, field.Selections, res)
}

func (ec *executionContext) _Element_error(ctx context.Context, field graphql.CollectedField, obj *models.Element) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Element",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Element().Error(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) _Element_mismatched(ctx context.Context, field graphql.CollectedField, obj *models.Element) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Element",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Element().Mismatched(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚕbool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_path(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Path(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*models.Element)
	fc.Result = res
	return ec.marshalOElement2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_date(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_date_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Date(rctx, args["filter"].(models.DateFilter))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_viewer(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Viewer(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*models.Viewer)
	fc.Result = res
	return ec.marshalOViewer2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐViewer(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_jsonEncoding(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().JSONEncoding(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_error(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_error_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Error(rctx, args["type"].(*models.ErrorType))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_complexity(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_complexity_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Complexity(rctx, args["value"].(int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query___type_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) _User_name(ctx context.Context, field graphql.CollectedField, obj *remote_api.User) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _User_likes(ctx context.Context, field graphql.CollectedField, obj *remote_api.User) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.User().Likes(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalNString2ᚕstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Viewer_user(ctx context.Context, field graphql.CollectedField, obj *models.Viewer) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Viewer",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.User, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*remote_api.User)
	fc.Result = res
	return ec.marshalOUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋremote_apiᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalN__DirectiveLocation2ᚕstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	fc.Result = res
	return ec.marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalN__TypeKind2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	fc.Result = res
	return ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	fc.Result = res
	return ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputDateFilter(ctx context.Context, obj interface{}) (models.DateFilter, error) {
	var it models.DateFilter
	var asMap = obj.(map[string]interface{})

	if _, present := asMap["timezone"]; !present {
		asMap["timezone"] = "UTC"
	}
	if _, present := asMap["op"]; !present {
		asMap["op"] = "EQ"
	}

	for k, v := range asMap {
		switch k {
		case "value":
			var err error
			it.Value, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "timezone":
			var err error
			it.Timezone, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "op":
			var err error
			it.Op, err = ec.unmarshalODATE_FILTER_OP2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilterOp(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var elementImplementors = []string{"Element"}

func (ec *executionContext) _Element(ctx context.Context, sel ast.SelectionSet, obj *models.Element) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, elementImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Element")
		case "child":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Element_child(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "error":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Element_error(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "mismatched":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Element_mismatched(ctx, field, obj)
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
		Stats:  graphql.FieldStats{Started: graphql.Now()},
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "path":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_path(ctx, field)
				return res
			})
		case "date":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_date(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "viewer":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_viewer(ctx, field)
				return res
			})
		case "jsonEncoding":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_jsonEncoding(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "error":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_error(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "complexity":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_complexity(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var userImplementors = []string{"User"}

func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *remote_api.User) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "name":
			out.Values[i] = ec._User_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "likes":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._User_likes(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var viewerImplementors = []string{"Viewer"}

func (ec *executionContext) _Viewer(ctx context.Context, sel ast.SelectionSet, obj *models.Viewer) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, viewerImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Viewer")
		case "user":
			out.Values[i] = ec._Viewer_user(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalNBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	res := graphql.MarshalBoolean(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNDateFilter2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilter(ctx context.Context, v interface{}) (models.DateFilter, error) {
	return ec.unmarshalInputDateFilter(ctx, v)
}

func (ec *executionContext) marshalNElement2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx context.Context, sel ast.SelectionSet, v models.Element) graphql.Marshaler {
	return ec._Element(ctx, sel, &v)
}

func (ec *executionContext) marshalNElement2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx context.Context, sel ast.SelectionSet, v *models.Element) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Element(ctx, sel, v)
}

func (ec *executionContext) unmarshalNInt2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalInt(v)
}

func (ec *executionContext) marshalNInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	res := graphql.MarshalInt(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNString2ᚕstring(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNString2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNString2ᚕstring(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNString2string(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v introspection.Directive) graphql.Marshaler {
	return ec.___Directive(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v []introspection.Directive) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalN__DirectiveLocation2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__DirectiveLocation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalN__DirectiveLocation2ᚕstring(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalN__DirectiveLocation2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalN__DirectiveLocation2ᚕstring(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__DirectiveLocation2string(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v introspection.EnumValue) graphql.Marshaler {
	return ec.___EnumValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v introspection.Field) graphql.Marshaler {
	return ec.___Field(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v introspection.InputValue) graphql.Marshaler {
	return ec.___InputValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func (ec *executionContext) unmarshalN__TypeKind2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__TypeKind2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
}

func (ec *executionContext) unmarshalOBoolean2ᚕbool(ctx context.Context, v interface{}) ([]bool, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]bool, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNBoolean2bool(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOBoolean2ᚕbool(ctx context.Context, sel ast.SelectionSet, v []bool) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNBoolean2bool(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) unmarshalOBoolean2ᚖbool(ctx context.Context, v interface{}) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOBoolean2bool(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOBoolean2ᚖbool(ctx context.Context, sel ast.SelectionSet, v *bool) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOBoolean2bool(ctx, sel, *v)
}

func (ec *executionContext) unmarshalODATE_FILTER_OP2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilterOp(ctx context.Context, v interface{}) (models.DateFilterOp, error) {
	var res models.DateFilterOp
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalODATE_FILTER_OP2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilterOp(ctx context.Context, sel ast.SelectionSet, v models.DateFilterOp) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalODATE_FILTER_OP2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilterOp(ctx context.Context, v interface{}) (*models.DateFilterOp, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalODATE_FILTER_OP2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilterOp(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalODATE_FILTER_OP2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐDateFilterOp(ctx context.Context, sel ast.SelectionSet, v *models.DateFilterOp) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) marshalOElement2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx context.Context, sel ast.SelectionSet, v models.Element) graphql.Marshaler {
	return ec._Element(ctx, sel, &v)
}

func (ec *executionContext) marshalOElement2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx context.Context, sel ast.SelectionSet, v []*models.Element) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOElement2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOElement2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐElement(ctx context.Context, sel ast.SelectionSet, v *models.Element) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Element(ctx, sel, v)
}

func (ec *executionContext) unmarshalOErrorType2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐErrorType(ctx context.Context, v interface{}) (models.ErrorType, error) {
	var res models.ErrorType
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalOErrorType2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐErrorType(ctx context.Context, sel ast.SelectionSet, v models.ErrorType) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOErrorType2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐErrorType(ctx context.Context, v interface{}) (*models.ErrorType, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOErrorType2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐErrorType(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOErrorType2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐErrorType(ctx context.Context, sel ast.SelectionSet, v *models.ErrorType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) unmarshalOInt2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalInt(v)
}

func (ec *executionContext) marshalOInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	return graphql.MarshalInt(v)
}

func (ec *executionContext) unmarshalOInt2ᚖint(ctx context.Context, v interface{}) (*int, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOInt2int(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOInt2ᚖint(ctx context.Context, sel ast.SelectionSet, v *int) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOInt2int(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalOString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOString2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOString2string(ctx, sel, *v)
}

func (ec *executionContext) marshalOUser2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋremote_apiᚐUser(ctx context.Context, sel ast.SelectionSet, v remote_api.User) graphql.Marshaler {
	return ec._User(ctx, sel, &v)
}

func (ec *executionContext) marshalOUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋremote_apiᚐUser(ctx context.Context, sel ast.SelectionSet, v *remote_api.User) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
}

func (ec *executionContext) marshalOViewer2githubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐViewer(ctx context.Context, sel ast.SelectionSet, v models.Viewer) graphql.Marshaler {
	return ec._Viewer(ctx, sel, &v)
}

func (ec *executionContext) marshalOViewer2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋintegrationᚋmodelsᚑgoᚐViewer(ctx context.Context, sel ast.SelectionSet, v *models.Viewer) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Viewer(ctx, sel, v)
}

func (ec *executionContext) marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v []introspection.EnumValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v []introspection.Field) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Schema2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v introspection.Schema) graphql.Marshaler {
	return ec.___Schema(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v *introspection.Schema) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, sel, v)
}

func (ec *executionContext) marshalO__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
