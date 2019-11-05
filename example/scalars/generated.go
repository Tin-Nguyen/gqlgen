package scalars

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/example/scalars/external"
	"github.com/99designs/gqlgen/example/scalars/model"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
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
	Query() QueryResolver
	User() UserResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Address struct {
		ID       func(childComplexity int) int
		Location func(childComplexity int) int
	}

	Query struct {
		Search func(childComplexity int, input *model.SearchArgs) int
		User   func(childComplexity int, id external.ObjectID) int
	}

	User struct {
		Address           func(childComplexity int) int
		Created           func(childComplexity int) int
		CustomResolver    func(childComplexity int) int
		ID                func(childComplexity int) int
		IsBanned          func(childComplexity int) int
		Name              func(childComplexity int) int
		PrimitiveResolver func(childComplexity int) int
		Tier              func(childComplexity int) int
	}
}

type QueryResolver interface {
	User(ctx context.Context, id external.ObjectID) (*model.User, error)
	Search(ctx context.Context, input *model.SearchArgs) ([]*model.User, error)
}
type UserResolver interface {
	PrimitiveResolver(ctx context.Context, obj *model.User) (string, error)
	CustomResolver(ctx context.Context, obj *model.User) (*model.Point, error)
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

	case "Address.id":
		if e.complexity.Address.ID == nil {
			break
		}

		return e.complexity.Address.ID(childComplexity), true

	case "Address.location":
		if e.complexity.Address.Location == nil {
			break
		}

		return e.complexity.Address.Location(childComplexity), true

	case "Query.search":
		if e.complexity.Query.Search == nil {
			break
		}

		args, err := ec.field_Query_search_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Search(childComplexity, args["input"].(*model.SearchArgs)), true

	case "Query.user":
		if e.complexity.Query.User == nil {
			break
		}

		args, err := ec.field_Query_user_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.User(childComplexity, args["id"].(external.ObjectID)), true

	case "User.address":
		if e.complexity.User.Address == nil {
			break
		}

		return e.complexity.User.Address(childComplexity), true

	case "User.created":
		if e.complexity.User.Created == nil {
			break
		}

		return e.complexity.User.Created(childComplexity), true

	case "User.customResolver":
		if e.complexity.User.CustomResolver == nil {
			break
		}

		return e.complexity.User.CustomResolver(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.isBanned":
		if e.complexity.User.IsBanned == nil {
			break
		}

		return e.complexity.User.IsBanned(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "User.primitiveResolver":
		if e.complexity.User.PrimitiveResolver == nil {
			break
		}

		return e.complexity.User.PrimitiveResolver(childComplexity), true

	case "User.tier":
		if e.complexity.User.Tier == nil {
			break
		}

		return e.complexity.User.Tier(childComplexity), true

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
	&ast.Source{Name: "schema.graphql", Input: `type Query {
    user(id: ID!): User
    search(input: SearchArgs = {location: "37,144", isBanned: false}): [User!]!
}

type User {
    id: ID!
    name: String!
    created: Timestamp
    isBanned: Banned!
    primitiveResolver: String!
    customResolver: Point!
    address: Address
    tier: Tier
}

type Address {
    id: ID!
    location: Point
}

input SearchArgs {
    location: Point
    createdAfter: Timestamp
    isBanned: Banned # TODO: This can be a Boolean again once multiple backing types are allowed
}

enum Tier {
    A
    B
    C
}

scalar Timestamp
scalar Point
scalar Banned
`},
)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

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

func (ec *executionContext) field_Query_search_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *model.SearchArgs
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = ec.unmarshalOSearchArgs2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐSearchArgs(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_user_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 external.ObjectID
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋexternalᚐObjectID(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
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

func (ec *executionContext) _Address_id(ctx context.Context, field graphql.CollectedField, obj *model.Address) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Address",
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
		return obj.ID, nil
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
	res := resTmp.(external.ObjectID)
	fc.Result = res
	return ec.marshalNID2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋexternalᚐObjectID(ctx, field.Selections, res)
}

func (ec *executionContext) _Address_location(ctx context.Context, field graphql.CollectedField, obj *model.Address) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Address",
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
		return obj.Location, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.Point)
	fc.Result = res
	return ec.marshalOPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_user(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_user_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().User(rctx, args["id"].(external.ObjectID))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.User)
	fc.Result = res
	return ec.marshalOUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_search(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_search_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Search(rctx, args["input"].(*model.SearchArgs))
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
	res := resTmp.([]*model.User)
	fc.Result = res
	return ec.marshalNUser2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx, field.Selections, res)
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

func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return obj.ID, nil
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
	res := resTmp.(external.ObjectID)
	fc.Result = res
	return ec.marshalNID2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋexternalᚐObjectID(ctx, field.Selections, res)
}

func (ec *executionContext) _User_name(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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

func (ec *executionContext) _User_created(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return obj.Created, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalOTimestamp2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) _User_isBanned(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return obj.IsBanned, nil
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
	res := resTmp.(model.Banned)
	fc.Result = res
	return ec.marshalNBanned2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐBanned(ctx, field.Selections, res)
}

func (ec *executionContext) _User_primitiveResolver(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return ec.resolvers.User().PrimitiveResolver(rctx, obj)
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

func (ec *executionContext) _User_customResolver(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return ec.resolvers.User().CustomResolver(rctx, obj)
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
	res := resTmp.(*model.Point)
	fc.Result = res
	return ec.marshalNPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx, field.Selections, res)
}

func (ec *executionContext) _User_address(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return obj.Address, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(model.Address)
	fc.Result = res
	return ec.marshalOAddress2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐAddress(ctx, field.Selections, res)
}

func (ec *executionContext) _User_tier(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
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
		return obj.Tier, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(model.Tier)
	fc.Result = res
	return ec.marshalOTier2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐTier(ctx, field.Selections, res)
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

func (ec *executionContext) unmarshalInputSearchArgs(ctx context.Context, obj interface{}) (model.SearchArgs, error) {
	var it model.SearchArgs
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "location":
			var err error
			it.Location, err = ec.unmarshalOPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx, v)
			if err != nil {
				return it, err
			}
		case "createdAfter":
			var err error
			it.CreatedAfter, err = ec.unmarshalOTimestamp2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
		case "isBanned":
			var err error
			it.IsBanned, err = ec.unmarshalOBanned2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐBanned(ctx, v)
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

var addressImplementors = []string{"Address"}

func (ec *executionContext) _Address(ctx context.Context, sel ast.SelectionSet, obj *model.Address) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, addressImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Address")
		case "id":
			out.Values[i] = ec._Address_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "location":
			out.Values[i] = ec._Address_location(ctx, field, obj)
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
		case "user":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_user(ctx, field)
				return res
			})
		case "search":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_search(ctx, field)
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

func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *model.User) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "id":
			out.Values[i] = ec._User_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "name":
			out.Values[i] = ec._User_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "created":
			out.Values[i] = ec._User_created(ctx, field, obj)
		case "isBanned":
			out.Values[i] = ec._User_isBanned(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "primitiveResolver":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._User_primitiveResolver(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "customResolver":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._User_customResolver(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "address":
			out.Values[i] = ec._User_address(ctx, field, obj)
		case "tier":
			out.Values[i] = ec._User_tier(ctx, field, obj)
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

func (ec *executionContext) unmarshalNBanned2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐBanned(ctx context.Context, v interface{}) (model.Banned, error) {
	var res model.Banned
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalNBanned2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐBanned(ctx context.Context, sel ast.SelectionSet, v model.Banned) graphql.Marshaler {
	return v
}

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

func (ec *executionContext) unmarshalNID2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋexternalᚐObjectID(ctx context.Context, v interface{}) (external.ObjectID, error) {
	return model.UnmarshalID(v)
}

func (ec *executionContext) marshalNID2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋexternalᚐObjectID(ctx context.Context, sel ast.SelectionSet, v external.ObjectID) graphql.Marshaler {
	res := model.MarshalID(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNPoint2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, v interface{}) (model.Point, error) {
	var res model.Point
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalNPoint2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, sel ast.SelectionSet, v model.Point) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, v interface{}) (*model.Point, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalNPoint2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalNPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, sel ast.SelectionSet, v *model.Point) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return v
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

func (ec *executionContext) marshalNUser2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v model.User) graphql.Marshaler {
	return ec._User(ctx, sel, &v)
}

func (ec *executionContext) marshalNUser2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v []*model.User) graphql.Marshaler {
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
			ret[i] = ec.marshalNUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx, sel, v[i])
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

func (ec *executionContext) marshalNUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v *model.User) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
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

func (ec *executionContext) marshalOAddress2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐAddress(ctx context.Context, sel ast.SelectionSet, v model.Address) graphql.Marshaler {
	return ec._Address(ctx, sel, &v)
}

func (ec *executionContext) unmarshalOBanned2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐBanned(ctx context.Context, v interface{}) (model.Banned, error) {
	var res model.Banned
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalOBanned2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐBanned(ctx context.Context, sel ast.SelectionSet, v model.Banned) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
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

func (ec *executionContext) unmarshalOPoint2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, v interface{}) (model.Point, error) {
	var res model.Point
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalOPoint2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, sel ast.SelectionSet, v model.Point) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, v interface{}) (*model.Point, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOPoint2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOPoint2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐPoint(ctx context.Context, sel ast.SelectionSet, v *model.Point) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) unmarshalOSearchArgs2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐSearchArgs(ctx context.Context, v interface{}) (model.SearchArgs, error) {
	return ec.unmarshalInputSearchArgs(ctx, v)
}

func (ec *executionContext) unmarshalOSearchArgs2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐSearchArgs(ctx context.Context, v interface{}) (*model.SearchArgs, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOSearchArgs2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐSearchArgs(ctx, v)
	return &res, err
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

func (ec *executionContext) unmarshalOTier2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐTier(ctx context.Context, v interface{}) (model.Tier, error) {
	var res model.Tier
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalOTier2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐTier(ctx context.Context, sel ast.SelectionSet, v model.Tier) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOTimestamp2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	return model.UnmarshalTimestamp(v)
}

func (ec *executionContext) marshalOTimestamp2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	return model.MarshalTimestamp(v)
}

func (ec *executionContext) unmarshalOTimestamp2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOTimestamp2timeᚐTime(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOTimestamp2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOTimestamp2timeᚐTime(ctx, sel, *v)
}

func (ec *executionContext) marshalOUser2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v model.User) graphql.Marshaler {
	return ec._User(ctx, sel, &v)
}

func (ec *executionContext) marshalOUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋscalarsᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v *model.User) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
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
