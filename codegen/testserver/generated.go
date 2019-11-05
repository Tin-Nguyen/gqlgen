package testserver

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	introspection1 "github.com/99designs/gqlgen/codegen/testserver/introspection"
	invalid_packagename "github.com/99designs/gqlgen/codegen/testserver/invalid-packagename"
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
	Errors() ErrorsResolver
	ForcedResolver() ForcedResolverResolver
	ModelMethods() ModelMethodsResolver
	OverlappingFields() OverlappingFieldsResolver
	Panics() PanicsResolver
	Primitive() PrimitiveResolver
	PrimitiveString() PrimitiveStringResolver
	Query() QueryResolver
	Subscription() SubscriptionResolver
	User() UserResolver
}

type DirectiveRoot struct {
	Custom func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

	Directive1 func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

	Directive2 func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

	Length func(ctx context.Context, obj interface{}, next graphql.Resolver, min int, max *int, message *string) (res interface{}, err error)

	Logged func(ctx context.Context, obj interface{}, next graphql.Resolver, id string) (res interface{}, err error)

	MakeNil func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

	Range func(ctx context.Context, obj interface{}, next graphql.Resolver, min *int, max *int) (res interface{}, err error)

	ToNull func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

	Unimplemented func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
}

type ComplexityRoot struct {
	A struct {
		ID func(childComplexity int) int
	}

	AIt struct {
		ID func(childComplexity int) int
	}

	AbIt struct {
		ID func(childComplexity int) int
	}

	Autobind struct {
		IdInt func(childComplexity int) int
		IdStr func(childComplexity int) int
		Int   func(childComplexity int) int
		Int32 func(childComplexity int) int
		Int64 func(childComplexity int) int
	}

	B struct {
		ID func(childComplexity int) int
	}

	Circle struct {
		Area   func(childComplexity int) int
		Radius func(childComplexity int) int
	}

	ContentPost struct {
		Foo func(childComplexity int) int
	}

	ContentUser struct {
		Foo func(childComplexity int) int
	}

	EmbeddedDefaultScalar struct {
		Value func(childComplexity int) int
	}

	EmbeddedPointer struct {
		ID    func(childComplexity int) int
		Title func(childComplexity int) int
	}

	Error struct {
		ErrorOnNonRequiredField func(childComplexity int) int
		ErrorOnRequiredField    func(childComplexity int) int
		ID                      func(childComplexity int) int
		NilOnRequiredField      func(childComplexity int) int
	}

	Errors struct {
		A func(childComplexity int) int
		B func(childComplexity int) int
		C func(childComplexity int) int
		D func(childComplexity int) int
		E func(childComplexity int) int
	}

	ForcedResolver struct {
		Field func(childComplexity int) int
	}

	InnerObject struct {
		ID func(childComplexity int) int
	}

	InvalidIdentifier struct {
		ID func(childComplexity int) int
	}

	It struct {
		ID func(childComplexity int) int
	}

	LoopA struct {
		B func(childComplexity int) int
	}

	LoopB struct {
		A func(childComplexity int) int
	}

	Map struct {
		ID func(childComplexity int) int
	}

	MapStringInterfaceType struct {
		A func(childComplexity int) int
		B func(childComplexity int) int
	}

	ModelMethods struct {
		NoContext     func(childComplexity int) int
		ResolverField func(childComplexity int) int
		WithContext   func(childComplexity int) int
	}

	ObjectDirectives struct {
		NullableText func(childComplexity int) int
		Text         func(childComplexity int) int
	}

	ObjectDirectivesWithCustomGoModel struct {
		NullableText func(childComplexity int) int
	}

	OuterObject struct {
		Inner func(childComplexity int) int
	}

	OverlappingFields struct {
		Foo    func(childComplexity int) int
		NewFoo func(childComplexity int) int
		OldFoo func(childComplexity int) int
	}

	Panics struct {
		ArgUnmarshal       func(childComplexity int, u []MarshalPanic) int
		FieldFuncMarshal   func(childComplexity int, u []MarshalPanic) int
		FieldScalarMarshal func(childComplexity int) int
	}

	Primitive struct {
		Squared func(childComplexity int) int
		Value   func(childComplexity int) int
	}

	PrimitiveString struct {
		Doubled func(childComplexity int) int
		Len     func(childComplexity int) int
		Value   func(childComplexity int) int
	}

	Query struct {
		Autobind                         func(childComplexity int) int
		Collision                        func(childComplexity int) int
		DefaultScalar                    func(childComplexity int, arg string) int
		DeprecatedField                  func(childComplexity int) int
		DirectiveArg                     func(childComplexity int, arg string) int
		DirectiveDouble                  func(childComplexity int) int
		DirectiveField                   func(childComplexity int) int
		DirectiveFieldDef                func(childComplexity int, ret string) int
		DirectiveInput                   func(childComplexity int, arg InputDirectives) int
		DirectiveInputNullable           func(childComplexity int, arg *InputDirectives) int
		DirectiveInputType               func(childComplexity int, arg InnerInput) int
		DirectiveNullableArg             func(childComplexity int, arg *int, arg2 *int, arg3 *string) int
		DirectiveObject                  func(childComplexity int) int
		DirectiveObjectWithCustomGoModel func(childComplexity int) int
		DirectiveUnimplemented           func(childComplexity int) int
		ErrorBubble                      func(childComplexity int) int
		Errors                           func(childComplexity int) int
		Fallback                         func(childComplexity int, arg FallbackToStringEncoding) int
		InputSlice                       func(childComplexity int, arg []string) int
		InvalidIdentifier                func(childComplexity int) int
		MapInput                         func(childComplexity int, input map[string]interface{}) int
		MapNestedStringInterface         func(childComplexity int, in *NestedMapInput) int
		MapStringInterface               func(childComplexity int, in map[string]interface{}) int
		ModelMethods                     func(childComplexity int) int
		NestedInputs                     func(childComplexity int, input [][]*OuterInput) int
		NestedOutputs                    func(childComplexity int) int
		NoShape                          func(childComplexity int) int
		NullableArg                      func(childComplexity int, arg *int) int
		OptionalUnion                    func(childComplexity int) int
		Overlapping                      func(childComplexity int) int
		Panics                           func(childComplexity int) int
		PrimitiveObject                  func(childComplexity int) int
		PrimitiveStringObject            func(childComplexity int) int
		Recursive                        func(childComplexity int, input *RecursiveInputSlice) int
		ScalarSlice                      func(childComplexity int) int
		ShapeUnion                       func(childComplexity int) int
		Shapes                           func(childComplexity int) int
		Slices                           func(childComplexity int) int
		User                             func(childComplexity int, id int) int
		Valid                            func(childComplexity int) int
		ValidType                        func(childComplexity int) int
		WrappedScalar                    func(childComplexity int) int
		WrappedStruct                    func(childComplexity int) int
	}

	Rectangle struct {
		Area   func(childComplexity int) int
		Length func(childComplexity int) int
		Width  func(childComplexity int) int
	}

	Slices struct {
		Test1 func(childComplexity int) int
		Test2 func(childComplexity int) int
		Test3 func(childComplexity int) int
		Test4 func(childComplexity int) int
	}

	Subscription struct {
		DirectiveArg           func(childComplexity int, arg string) int
		DirectiveDouble        func(childComplexity int) int
		DirectiveNullableArg   func(childComplexity int, arg *int, arg2 *int, arg3 *string) int
		DirectiveUnimplemented func(childComplexity int) int
		InitPayload            func(childComplexity int) int
		Updated                func(childComplexity int) int
	}

	User struct {
		Created func(childComplexity int) int
		Friends func(childComplexity int) int
		ID      func(childComplexity int) int
		Updated func(childComplexity int) int
	}

	ValidType struct {
		DifferentCase      func(childComplexity int) int
		DifferentCaseOld   func(childComplexity int) int
		ValidArgs          func(childComplexity int, breakArg string, defaultArg string, funcArg string, interfaceArg string, selectArg string, caseArg string, deferArg string, goArg string, mapArg string, structArg string, chanArg string, elseArg string, gotoArg string, packageArg string, switchArg string, constArg string, fallthroughArg string, ifArg string, rangeArg string, typeArg string, continueArg string, forArg string, importArg string, returnArg string, varArg string, _Arg string) int
		ValidInputKeywords func(childComplexity int, input *ValidInput) int
	}

	WrappedStruct struct {
		Name func(childComplexity int) int
	}

	XXIt struct {
		ID func(childComplexity int) int
	}

	XxIt struct {
		ID func(childComplexity int) int
	}

	AsdfIt struct {
		ID func(childComplexity int) int
	}

	IIt struct {
		ID func(childComplexity int) int
	}
}

type ErrorsResolver interface {
	A(ctx context.Context, obj *Errors) (*Error, error)
	B(ctx context.Context, obj *Errors) (*Error, error)
	C(ctx context.Context, obj *Errors) (*Error, error)
	D(ctx context.Context, obj *Errors) (*Error, error)
	E(ctx context.Context, obj *Errors) (*Error, error)
}
type ForcedResolverResolver interface {
	Field(ctx context.Context, obj *ForcedResolver) (*Circle, error)
}
type ModelMethodsResolver interface {
	ResolverField(ctx context.Context, obj *ModelMethods) (bool, error)
}
type OverlappingFieldsResolver interface {
	OldFoo(ctx context.Context, obj *OverlappingFields) (int, error)
}
type PanicsResolver interface {
	FieldScalarMarshal(ctx context.Context, obj *Panics) ([]MarshalPanic, error)

	ArgUnmarshal(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error)
}
type PrimitiveResolver interface {
	Value(ctx context.Context, obj *Primitive) (int, error)
}
type PrimitiveStringResolver interface {
	Value(ctx context.Context, obj *PrimitiveString) (string, error)

	Len(ctx context.Context, obj *PrimitiveString) (int, error)
}
type QueryResolver interface {
	InvalidIdentifier(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error)
	Collision(ctx context.Context) (*introspection1.It, error)
	MapInput(ctx context.Context, input map[string]interface{}) (*bool, error)
	Recursive(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
	NestedInputs(ctx context.Context, input [][]*OuterInput) (*bool, error)
	NestedOutputs(ctx context.Context) ([][]*OuterObject, error)
	ModelMethods(ctx context.Context) (*ModelMethods, error)
	User(ctx context.Context, id int) (*User, error)
	NullableArg(ctx context.Context, arg *int) (*string, error)
	InputSlice(ctx context.Context, arg []string) (bool, error)
	ShapeUnion(ctx context.Context) (ShapeUnion, error)
	Autobind(ctx context.Context) (*Autobind, error)
	DeprecatedField(ctx context.Context) (string, error)
	Overlapping(ctx context.Context) (*OverlappingFields, error)
	DirectiveArg(ctx context.Context, arg string) (*string, error)
	DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int, arg3 *string) (*string, error)
	DirectiveInputNullable(ctx context.Context, arg *InputDirectives) (*string, error)
	DirectiveInput(ctx context.Context, arg InputDirectives) (*string, error)
	DirectiveInputType(ctx context.Context, arg InnerInput) (*string, error)
	DirectiveObject(ctx context.Context) (*ObjectDirectives, error)
	DirectiveObjectWithCustomGoModel(ctx context.Context) (*ObjectDirectivesWithCustomGoModel, error)
	DirectiveFieldDef(ctx context.Context, ret string) (string, error)
	DirectiveField(ctx context.Context) (*string, error)
	DirectiveDouble(ctx context.Context) (*string, error)
	DirectiveUnimplemented(ctx context.Context) (*string, error)
	Shapes(ctx context.Context) ([]Shape, error)
	NoShape(ctx context.Context) (Shape, error)
	MapStringInterface(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error)
	MapNestedStringInterface(ctx context.Context, in *NestedMapInput) (map[string]interface{}, error)
	ErrorBubble(ctx context.Context) (*Error, error)
	Errors(ctx context.Context) (*Errors, error)
	Valid(ctx context.Context) (string, error)
	Panics(ctx context.Context) (*Panics, error)
	PrimitiveObject(ctx context.Context) ([]Primitive, error)
	PrimitiveStringObject(ctx context.Context) ([]PrimitiveString, error)
	DefaultScalar(ctx context.Context, arg string) (string, error)
	Slices(ctx context.Context) (*Slices, error)
	ScalarSlice(ctx context.Context) ([]byte, error)
	Fallback(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error)
	OptionalUnion(ctx context.Context) (TestUnion, error)
	ValidType(ctx context.Context) (*ValidType, error)
	WrappedStruct(ctx context.Context) (*WrappedStruct, error)
	WrappedScalar(ctx context.Context) (WrappedScalar, error)
}
type SubscriptionResolver interface {
	Updated(ctx context.Context) (<-chan string, error)
	InitPayload(ctx context.Context) (<-chan string, error)
	DirectiveArg(ctx context.Context, arg string) (<-chan *string, error)
	DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int, arg3 *string) (<-chan *string, error)
	DirectiveDouble(ctx context.Context) (<-chan *string, error)
	DirectiveUnimplemented(ctx context.Context) (<-chan *string, error)
}
type UserResolver interface {
	Friends(ctx context.Context, obj *User) ([]*User, error)
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

	case "A.id":
		if e.complexity.A.ID == nil {
			break
		}

		return e.complexity.A.ID(childComplexity), true

	case "AIt.id":
		if e.complexity.AIt.ID == nil {
			break
		}

		return e.complexity.AIt.ID(childComplexity), true

	case "AbIt.id":
		if e.complexity.AbIt.ID == nil {
			break
		}

		return e.complexity.AbIt.ID(childComplexity), true

	case "Autobind.idInt":
		if e.complexity.Autobind.IdInt == nil {
			break
		}

		return e.complexity.Autobind.IdInt(childComplexity), true

	case "Autobind.idStr":
		if e.complexity.Autobind.IdStr == nil {
			break
		}

		return e.complexity.Autobind.IdStr(childComplexity), true

	case "Autobind.int":
		if e.complexity.Autobind.Int == nil {
			break
		}

		return e.complexity.Autobind.Int(childComplexity), true

	case "Autobind.int32":
		if e.complexity.Autobind.Int32 == nil {
			break
		}

		return e.complexity.Autobind.Int32(childComplexity), true

	case "Autobind.int64":
		if e.complexity.Autobind.Int64 == nil {
			break
		}

		return e.complexity.Autobind.Int64(childComplexity), true

	case "B.id":
		if e.complexity.B.ID == nil {
			break
		}

		return e.complexity.B.ID(childComplexity), true

	case "Circle.area":
		if e.complexity.Circle.Area == nil {
			break
		}

		return e.complexity.Circle.Area(childComplexity), true

	case "Circle.radius":
		if e.complexity.Circle.Radius == nil {
			break
		}

		return e.complexity.Circle.Radius(childComplexity), true

	case "Content_Post.foo":
		if e.complexity.ContentPost.Foo == nil {
			break
		}

		return e.complexity.ContentPost.Foo(childComplexity), true

	case "Content_User.foo":
		if e.complexity.ContentUser.Foo == nil {
			break
		}

		return e.complexity.ContentUser.Foo(childComplexity), true

	case "EmbeddedDefaultScalar.value":
		if e.complexity.EmbeddedDefaultScalar.Value == nil {
			break
		}

		return e.complexity.EmbeddedDefaultScalar.Value(childComplexity), true

	case "EmbeddedPointer.ID":
		if e.complexity.EmbeddedPointer.ID == nil {
			break
		}

		return e.complexity.EmbeddedPointer.ID(childComplexity), true

	case "EmbeddedPointer.Title":
		if e.complexity.EmbeddedPointer.Title == nil {
			break
		}

		return e.complexity.EmbeddedPointer.Title(childComplexity), true

	case "Error.errorOnNonRequiredField":
		if e.complexity.Error.ErrorOnNonRequiredField == nil {
			break
		}

		return e.complexity.Error.ErrorOnNonRequiredField(childComplexity), true

	case "Error.errorOnRequiredField":
		if e.complexity.Error.ErrorOnRequiredField == nil {
			break
		}

		return e.complexity.Error.ErrorOnRequiredField(childComplexity), true

	case "Error.id":
		if e.complexity.Error.ID == nil {
			break
		}

		return e.complexity.Error.ID(childComplexity), true

	case "Error.nilOnRequiredField":
		if e.complexity.Error.NilOnRequiredField == nil {
			break
		}

		return e.complexity.Error.NilOnRequiredField(childComplexity), true

	case "Errors.a":
		if e.complexity.Errors.A == nil {
			break
		}

		return e.complexity.Errors.A(childComplexity), true

	case "Errors.b":
		if e.complexity.Errors.B == nil {
			break
		}

		return e.complexity.Errors.B(childComplexity), true

	case "Errors.c":
		if e.complexity.Errors.C == nil {
			break
		}

		return e.complexity.Errors.C(childComplexity), true

	case "Errors.d":
		if e.complexity.Errors.D == nil {
			break
		}

		return e.complexity.Errors.D(childComplexity), true

	case "Errors.e":
		if e.complexity.Errors.E == nil {
			break
		}

		return e.complexity.Errors.E(childComplexity), true

	case "ForcedResolver.field":
		if e.complexity.ForcedResolver.Field == nil {
			break
		}

		return e.complexity.ForcedResolver.Field(childComplexity), true

	case "InnerObject.id":
		if e.complexity.InnerObject.ID == nil {
			break
		}

		return e.complexity.InnerObject.ID(childComplexity), true

	case "InvalidIdentifier.id":
		if e.complexity.InvalidIdentifier.ID == nil {
			break
		}

		return e.complexity.InvalidIdentifier.ID(childComplexity), true

	case "It.id":
		if e.complexity.It.ID == nil {
			break
		}

		return e.complexity.It.ID(childComplexity), true

	case "LoopA.b":
		if e.complexity.LoopA.B == nil {
			break
		}

		return e.complexity.LoopA.B(childComplexity), true

	case "LoopB.a":
		if e.complexity.LoopB.A == nil {
			break
		}

		return e.complexity.LoopB.A(childComplexity), true

	case "Map.id":
		if e.complexity.Map.ID == nil {
			break
		}

		return e.complexity.Map.ID(childComplexity), true

	case "MapStringInterfaceType.a":
		if e.complexity.MapStringInterfaceType.A == nil {
			break
		}

		return e.complexity.MapStringInterfaceType.A(childComplexity), true

	case "MapStringInterfaceType.b":
		if e.complexity.MapStringInterfaceType.B == nil {
			break
		}

		return e.complexity.MapStringInterfaceType.B(childComplexity), true

	case "ModelMethods.noContext":
		if e.complexity.ModelMethods.NoContext == nil {
			break
		}

		return e.complexity.ModelMethods.NoContext(childComplexity), true

	case "ModelMethods.resolverField":
		if e.complexity.ModelMethods.ResolverField == nil {
			break
		}

		return e.complexity.ModelMethods.ResolverField(childComplexity), true

	case "ModelMethods.withContext":
		if e.complexity.ModelMethods.WithContext == nil {
			break
		}

		return e.complexity.ModelMethods.WithContext(childComplexity), true

	case "ObjectDirectives.nullableText":
		if e.complexity.ObjectDirectives.NullableText == nil {
			break
		}

		return e.complexity.ObjectDirectives.NullableText(childComplexity), true

	case "ObjectDirectives.text":
		if e.complexity.ObjectDirectives.Text == nil {
			break
		}

		return e.complexity.ObjectDirectives.Text(childComplexity), true

	case "ObjectDirectivesWithCustomGoModel.nullableText":
		if e.complexity.ObjectDirectivesWithCustomGoModel.NullableText == nil {
			break
		}

		return e.complexity.ObjectDirectivesWithCustomGoModel.NullableText(childComplexity), true

	case "OuterObject.inner":
		if e.complexity.OuterObject.Inner == nil {
			break
		}

		return e.complexity.OuterObject.Inner(childComplexity), true

	case "OverlappingFields.oneFoo", "OverlappingFields.twoFoo":
		if e.complexity.OverlappingFields.Foo == nil {
			break
		}

		return e.complexity.OverlappingFields.Foo(childComplexity), true

	case "OverlappingFields.newFoo", "OverlappingFields.new_foo":
		if e.complexity.OverlappingFields.NewFoo == nil {
			break
		}

		return e.complexity.OverlappingFields.NewFoo(childComplexity), true

	case "OverlappingFields.oldFoo":
		if e.complexity.OverlappingFields.OldFoo == nil {
			break
		}

		return e.complexity.OverlappingFields.OldFoo(childComplexity), true

	case "Panics.argUnmarshal":
		if e.complexity.Panics.ArgUnmarshal == nil {
			break
		}

		args, err := ec.field_Panics_argUnmarshal_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Panics.ArgUnmarshal(childComplexity, args["u"].([]MarshalPanic)), true

	case "Panics.fieldFuncMarshal":
		if e.complexity.Panics.FieldFuncMarshal == nil {
			break
		}

		args, err := ec.field_Panics_fieldFuncMarshal_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Panics.FieldFuncMarshal(childComplexity, args["u"].([]MarshalPanic)), true

	case "Panics.fieldScalarMarshal":
		if e.complexity.Panics.FieldScalarMarshal == nil {
			break
		}

		return e.complexity.Panics.FieldScalarMarshal(childComplexity), true

	case "Primitive.squared":
		if e.complexity.Primitive.Squared == nil {
			break
		}

		return e.complexity.Primitive.Squared(childComplexity), true

	case "Primitive.value":
		if e.complexity.Primitive.Value == nil {
			break
		}

		return e.complexity.Primitive.Value(childComplexity), true

	case "PrimitiveString.doubled":
		if e.complexity.PrimitiveString.Doubled == nil {
			break
		}

		return e.complexity.PrimitiveString.Doubled(childComplexity), true

	case "PrimitiveString.len":
		if e.complexity.PrimitiveString.Len == nil {
			break
		}

		return e.complexity.PrimitiveString.Len(childComplexity), true

	case "PrimitiveString.value":
		if e.complexity.PrimitiveString.Value == nil {
			break
		}

		return e.complexity.PrimitiveString.Value(childComplexity), true

	case "Query.autobind":
		if e.complexity.Query.Autobind == nil {
			break
		}

		return e.complexity.Query.Autobind(childComplexity), true

	case "Query.collision":
		if e.complexity.Query.Collision == nil {
			break
		}

		return e.complexity.Query.Collision(childComplexity), true

	case "Query.defaultScalar":
		if e.complexity.Query.DefaultScalar == nil {
			break
		}

		args, err := ec.field_Query_defaultScalar_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DefaultScalar(childComplexity, args["arg"].(string)), true

	case "Query.deprecatedField":
		if e.complexity.Query.DeprecatedField == nil {
			break
		}

		return e.complexity.Query.DeprecatedField(childComplexity), true

	case "Query.directiveArg":
		if e.complexity.Query.DirectiveArg == nil {
			break
		}

		args, err := ec.field_Query_directiveArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveArg(childComplexity, args["arg"].(string)), true

	case "Query.directiveDouble":
		if e.complexity.Query.DirectiveDouble == nil {
			break
		}

		return e.complexity.Query.DirectiveDouble(childComplexity), true

	case "Query.directiveField":
		if e.complexity.Query.DirectiveField == nil {
			break
		}

		return e.complexity.Query.DirectiveField(childComplexity), true

	case "Query.directiveFieldDef":
		if e.complexity.Query.DirectiveFieldDef == nil {
			break
		}

		args, err := ec.field_Query_directiveFieldDef_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveFieldDef(childComplexity, args["ret"].(string)), true

	case "Query.directiveInput":
		if e.complexity.Query.DirectiveInput == nil {
			break
		}

		args, err := ec.field_Query_directiveInput_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveInput(childComplexity, args["arg"].(InputDirectives)), true

	case "Query.directiveInputNullable":
		if e.complexity.Query.DirectiveInputNullable == nil {
			break
		}

		args, err := ec.field_Query_directiveInputNullable_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveInputNullable(childComplexity, args["arg"].(*InputDirectives)), true

	case "Query.directiveInputType":
		if e.complexity.Query.DirectiveInputType == nil {
			break
		}

		args, err := ec.field_Query_directiveInputType_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveInputType(childComplexity, args["arg"].(InnerInput)), true

	case "Query.directiveNullableArg":
		if e.complexity.Query.DirectiveNullableArg == nil {
			break
		}

		args, err := ec.field_Query_directiveNullableArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveNullableArg(childComplexity, args["arg"].(*int), args["arg2"].(*int), args["arg3"].(*string)), true

	case "Query.directiveObject":
		if e.complexity.Query.DirectiveObject == nil {
			break
		}

		return e.complexity.Query.DirectiveObject(childComplexity), true

	case "Query.directiveObjectWithCustomGoModel":
		if e.complexity.Query.DirectiveObjectWithCustomGoModel == nil {
			break
		}

		return e.complexity.Query.DirectiveObjectWithCustomGoModel(childComplexity), true

	case "Query.directiveUnimplemented":
		if e.complexity.Query.DirectiveUnimplemented == nil {
			break
		}

		return e.complexity.Query.DirectiveUnimplemented(childComplexity), true

	case "Query.errorBubble":
		if e.complexity.Query.ErrorBubble == nil {
			break
		}

		return e.complexity.Query.ErrorBubble(childComplexity), true

	case "Query.errors":
		if e.complexity.Query.Errors == nil {
			break
		}

		return e.complexity.Query.Errors(childComplexity), true

	case "Query.fallback":
		if e.complexity.Query.Fallback == nil {
			break
		}

		args, err := ec.field_Query_fallback_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Fallback(childComplexity, args["arg"].(FallbackToStringEncoding)), true

	case "Query.inputSlice":
		if e.complexity.Query.InputSlice == nil {
			break
		}

		args, err := ec.field_Query_inputSlice_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.InputSlice(childComplexity, args["arg"].([]string)), true

	case "Query.invalidIdentifier":
		if e.complexity.Query.InvalidIdentifier == nil {
			break
		}

		return e.complexity.Query.InvalidIdentifier(childComplexity), true

	case "Query.mapInput":
		if e.complexity.Query.MapInput == nil {
			break
		}

		args, err := ec.field_Query_mapInput_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.MapInput(childComplexity, args["input"].(map[string]interface{})), true

	case "Query.mapNestedStringInterface":
		if e.complexity.Query.MapNestedStringInterface == nil {
			break
		}

		args, err := ec.field_Query_mapNestedStringInterface_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.MapNestedStringInterface(childComplexity, args["in"].(*NestedMapInput)), true

	case "Query.mapStringInterface":
		if e.complexity.Query.MapStringInterface == nil {
			break
		}

		args, err := ec.field_Query_mapStringInterface_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.MapStringInterface(childComplexity, args["in"].(map[string]interface{})), true

	case "Query.modelMethods":
		if e.complexity.Query.ModelMethods == nil {
			break
		}

		return e.complexity.Query.ModelMethods(childComplexity), true

	case "Query.nestedInputs":
		if e.complexity.Query.NestedInputs == nil {
			break
		}

		args, err := ec.field_Query_nestedInputs_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.NestedInputs(childComplexity, args["input"].([][]*OuterInput)), true

	case "Query.nestedOutputs":
		if e.complexity.Query.NestedOutputs == nil {
			break
		}

		return e.complexity.Query.NestedOutputs(childComplexity), true

	case "Query.noShape":
		if e.complexity.Query.NoShape == nil {
			break
		}

		return e.complexity.Query.NoShape(childComplexity), true

	case "Query.nullableArg":
		if e.complexity.Query.NullableArg == nil {
			break
		}

		args, err := ec.field_Query_nullableArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.NullableArg(childComplexity, args["arg"].(*int)), true

	case "Query.optionalUnion":
		if e.complexity.Query.OptionalUnion == nil {
			break
		}

		return e.complexity.Query.OptionalUnion(childComplexity), true

	case "Query.overlapping":
		if e.complexity.Query.Overlapping == nil {
			break
		}

		return e.complexity.Query.Overlapping(childComplexity), true

	case "Query.panics":
		if e.complexity.Query.Panics == nil {
			break
		}

		return e.complexity.Query.Panics(childComplexity), true

	case "Query.primitiveObject":
		if e.complexity.Query.PrimitiveObject == nil {
			break
		}

		return e.complexity.Query.PrimitiveObject(childComplexity), true

	case "Query.primitiveStringObject":
		if e.complexity.Query.PrimitiveStringObject == nil {
			break
		}

		return e.complexity.Query.PrimitiveStringObject(childComplexity), true

	case "Query.recursive":
		if e.complexity.Query.Recursive == nil {
			break
		}

		args, err := ec.field_Query_recursive_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Recursive(childComplexity, args["input"].(*RecursiveInputSlice)), true

	case "Query.scalarSlice":
		if e.complexity.Query.ScalarSlice == nil {
			break
		}

		return e.complexity.Query.ScalarSlice(childComplexity), true

	case "Query.shapeUnion":
		if e.complexity.Query.ShapeUnion == nil {
			break
		}

		return e.complexity.Query.ShapeUnion(childComplexity), true

	case "Query.shapes":
		if e.complexity.Query.Shapes == nil {
			break
		}

		return e.complexity.Query.Shapes(childComplexity), true

	case "Query.slices":
		if e.complexity.Query.Slices == nil {
			break
		}

		return e.complexity.Query.Slices(childComplexity), true

	case "Query.user":
		if e.complexity.Query.User == nil {
			break
		}

		args, err := ec.field_Query_user_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.User(childComplexity, args["id"].(int)), true

	case "Query.valid":
		if e.complexity.Query.Valid == nil {
			break
		}

		return e.complexity.Query.Valid(childComplexity), true

	case "Query.validType":
		if e.complexity.Query.ValidType == nil {
			break
		}

		return e.complexity.Query.ValidType(childComplexity), true

	case "Query.wrappedScalar":
		if e.complexity.Query.WrappedScalar == nil {
			break
		}

		return e.complexity.Query.WrappedScalar(childComplexity), true

	case "Query.wrappedStruct":
		if e.complexity.Query.WrappedStruct == nil {
			break
		}

		return e.complexity.Query.WrappedStruct(childComplexity), true

	case "Rectangle.area":
		if e.complexity.Rectangle.Area == nil {
			break
		}

		return e.complexity.Rectangle.Area(childComplexity), true

	case "Rectangle.length":
		if e.complexity.Rectangle.Length == nil {
			break
		}

		return e.complexity.Rectangle.Length(childComplexity), true

	case "Rectangle.width":
		if e.complexity.Rectangle.Width == nil {
			break
		}

		return e.complexity.Rectangle.Width(childComplexity), true

	case "Slices.test1":
		if e.complexity.Slices.Test1 == nil {
			break
		}

		return e.complexity.Slices.Test1(childComplexity), true

	case "Slices.test2":
		if e.complexity.Slices.Test2 == nil {
			break
		}

		return e.complexity.Slices.Test2(childComplexity), true

	case "Slices.test3":
		if e.complexity.Slices.Test3 == nil {
			break
		}

		return e.complexity.Slices.Test3(childComplexity), true

	case "Slices.test4":
		if e.complexity.Slices.Test4 == nil {
			break
		}

		return e.complexity.Slices.Test4(childComplexity), true

	case "Subscription.directiveArg":
		if e.complexity.Subscription.DirectiveArg == nil {
			break
		}

		args, err := ec.field_Subscription_directiveArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.DirectiveArg(childComplexity, args["arg"].(string)), true

	case "Subscription.directiveDouble":
		if e.complexity.Subscription.DirectiveDouble == nil {
			break
		}

		return e.complexity.Subscription.DirectiveDouble(childComplexity), true

	case "Subscription.directiveNullableArg":
		if e.complexity.Subscription.DirectiveNullableArg == nil {
			break
		}

		args, err := ec.field_Subscription_directiveNullableArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.DirectiveNullableArg(childComplexity, args["arg"].(*int), args["arg2"].(*int), args["arg3"].(*string)), true

	case "Subscription.directiveUnimplemented":
		if e.complexity.Subscription.DirectiveUnimplemented == nil {
			break
		}

		return e.complexity.Subscription.DirectiveUnimplemented(childComplexity), true

	case "Subscription.initPayload":
		if e.complexity.Subscription.InitPayload == nil {
			break
		}

		return e.complexity.Subscription.InitPayload(childComplexity), true

	case "Subscription.updated":
		if e.complexity.Subscription.Updated == nil {
			break
		}

		return e.complexity.Subscription.Updated(childComplexity), true

	case "User.created":
		if e.complexity.User.Created == nil {
			break
		}

		return e.complexity.User.Created(childComplexity), true

	case "User.friends":
		if e.complexity.User.Friends == nil {
			break
		}

		return e.complexity.User.Friends(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.updated":
		if e.complexity.User.Updated == nil {
			break
		}

		return e.complexity.User.Updated(childComplexity), true

	case "ValidType.differentCase":
		if e.complexity.ValidType.DifferentCase == nil {
			break
		}

		return e.complexity.ValidType.DifferentCase(childComplexity), true

	case "ValidType.different_case":
		if e.complexity.ValidType.DifferentCaseOld == nil {
			break
		}

		return e.complexity.ValidType.DifferentCaseOld(childComplexity), true

	case "ValidType.validArgs":
		if e.complexity.ValidType.ValidArgs == nil {
			break
		}

		args, err := ec.field_ValidType_validArgs_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.ValidType.ValidArgs(childComplexity, args["break"].(string), args["default"].(string), args["func"].(string), args["interface"].(string), args["select"].(string), args["case"].(string), args["defer"].(string), args["go"].(string), args["map"].(string), args["struct"].(string), args["chan"].(string), args["else"].(string), args["goto"].(string), args["package"].(string), args["switch"].(string), args["const"].(string), args["fallthrough"].(string), args["if"].(string), args["range"].(string), args["type"].(string), args["continue"].(string), args["for"].(string), args["import"].(string), args["return"].(string), args["var"].(string), args["_"].(string)), true

	case "ValidType.validInputKeywords":
		if e.complexity.ValidType.ValidInputKeywords == nil {
			break
		}

		args, err := ec.field_ValidType_validInputKeywords_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.ValidType.ValidInputKeywords(childComplexity, args["input"].(*ValidInput)), true

	case "WrappedStruct.name":
		if e.complexity.WrappedStruct.Name == nil {
			break
		}

		return e.complexity.WrappedStruct.Name(childComplexity), true

	case "XXIt.id":
		if e.complexity.XXIt.ID == nil {
			break
		}

		return e.complexity.XXIt.ID(childComplexity), true

	case "XxIt.id":
		if e.complexity.XxIt.ID == nil {
			break
		}

		return e.complexity.XxIt.ID(childComplexity), true

	case "asdfIt.id":
		if e.complexity.AsdfIt.ID == nil {
			break
		}

		return e.complexity.AsdfIt.ID(childComplexity), true

	case "iIt.id":
		if e.complexity.IIt.ID == nil {
			break
		}

		return e.complexity.IIt.ID(childComplexity), true

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
	case ast.Subscription:
		next := ec._Subscription(ctx, rc.Operation.SelectionSet)

		var buf bytes.Buffer
		return func(ctx context.Context) *graphql.Response {
			buf.Reset()
			data := next()

			if data == nil {
				return nil
			}
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
	&ast.Source{Name: "builtinscalar.graphql", Input: `
"""
Since gqlgen defines default implementation for a Map scalar, this tests that the builtin is _not_
added to the TypeMap
"""
type Map {
    id: ID!
}
`},
	&ast.Source{Name: "complexity.graphql", Input: `extend type Query {
    overlapping: OverlappingFields
}

type OverlappingFields {
  oneFoo: Int! @goField(name: "foo")
  twoFoo: Int! @goField(name: "foo")
  oldFoo: Int! @goField(name: "foo", forceResolver: true)
  newFoo: Int!
  new_foo: Int!
}
`},
	&ast.Source{Name: "directive.graphql", Input: `directive @length(min: Int!, max: Int, message: String) on ARGUMENT_DEFINITION | INPUT_FIELD_DEFINITION | FIELD_DEFINITION
directive @range(min: Int = 0, max: Int) on ARGUMENT_DEFINITION
directive @custom on ARGUMENT_DEFINITION
directive @logged(id: UUID!) on FIELD
directive @toNull on ARGUMENT_DEFINITION | INPUT_FIELD_DEFINITION | FIELD_DEFINITION
directive @directive1 on FIELD_DEFINITION
directive @directive2 on FIELD_DEFINITION
directive @unimplemented on FIELD_DEFINITION

extend type Query {
    directiveArg(arg: String! @length(min:1, max: 255, message: "invalid length")): String
    directiveNullableArg(arg: Int @range(min:0), arg2: Int @range, arg3: String @toNull): String
    directiveInputNullable(arg: InputDirectives): String
    directiveInput(arg: InputDirectives!): String
    directiveInputType(arg: InnerInput! @custom): String
    directiveObject: ObjectDirectives
    directiveObjectWithCustomGoModel: ObjectDirectivesWithCustomGoModel
    directiveFieldDef(ret: String!): String! @length(min: 1, message: "not valid")
    directiveField: String
    directiveDouble: String @directive1 @directive2
    directiveUnimplemented: String @unimplemented
}

extend type Subscription {
    directiveArg(arg: String! @length(min:1, max: 255, message: "invalid length")): String
    directiveNullableArg(arg: Int @range(min:0), arg2: Int @range, arg3: String @toNull): String
    directiveDouble: String @directive1 @directive2
    directiveUnimplemented: String @unimplemented
}

input InputDirectives {
    text: String! @length(min: 0, max: 7, message: "not valid")
    nullableText: String @toNull
    inner: InnerDirectives!
    innerNullable: InnerDirectives
    thirdParty: ThirdParty @length(min: 0, max: 7)
}

input InnerDirectives {
    message: String! @length(min: 1, message: "not valid")
}

type ObjectDirectives {
    text: String! @length(min: 0, max: 7, message: "not valid")
    nullableText: String @toNull
}

type ObjectDirectivesWithCustomGoModel {
    nullableText: String @toNull
}
`},
	&ast.Source{Name: "interfaces.graphql", Input: `extend type Query {
    shapes: [Shape]
    noShape: Shape @makeNil
}

interface Shape {
    area: Float
}
type Circle implements Shape {
    radius: Float
    area: Float
}
type Rectangle implements Shape {
    length: Float
    width: Float
    area: Float
}
union ShapeUnion @goModel(model:"testserver.ShapeUnion") = Circle | Rectangle

directive @makeNil on FIELD_DEFINITION
`},
	&ast.Source{Name: "loops.graphql", Input: `type LoopA {
    b: LoopB!
}

type LoopB {
    a: LoopA!
}
`},
	&ast.Source{Name: "maps.graphql", Input: `extend type Query {
    mapStringInterface(in: MapStringInterfaceInput): MapStringInterfaceType
    mapNestedStringInterface(in: NestedMapInput): MapStringInterfaceType
}

type MapStringInterfaceType @goModel(model: "map[string]interface{}") {
    a: String
    b: Int
}

input MapStringInterfaceInput @goModel(model: "map[string]interface{}") {
    a: String
    b: Int
}

input NestedMapInput {
    map: MapStringInterfaceInput
}
`},
	&ast.Source{Name: "nulls.graphql", Input: `extend type Query {
    errorBubble: Error
    errors: Errors
    valid: String!
}

type Errors {
    a: Error!
    b: Error!
    c: Error!
    d: Error!
    e: Error!
}

type Error {
    id: ID!
    errorOnNonRequiredField: String
    errorOnRequiredField: String!
    nilOnRequiredField: String!
}
`},
	&ast.Source{Name: "panics.graphql", Input: `extend type Query {
    panics: Panics
}

type Panics {
    fieldScalarMarshal: [MarshalPanic!]!
    fieldFuncMarshal(u: [MarshalPanic!]!): [MarshalPanic!]!
    argUnmarshal(u: [MarshalPanic!]!): Boolean!

}

scalar MarshalPanic
`},
	&ast.Source{Name: "primitive_objects.graphql", Input: `extend type Query {
    primitiveObject: [Primitive!]!
    primitiveStringObject: [PrimitiveString!]!
}

type Primitive {
    value: Int!
    squared: Int!
}

type PrimitiveString {
    value: String!
    doubled: String!
    len: Int!
}
`},
	&ast.Source{Name: "scalar_default.graphql", Input: `extend type Query {
    defaultScalar(arg: DefaultScalarImplementation! = "default"): DefaultScalarImplementation!
}

""" This doesnt have an implementation in the typemap, so it should act like a string """
scalar DefaultScalarImplementation

type EmbeddedDefaultScalar {
    value: DefaultScalarImplementation
}
`},
	&ast.Source{Name: "schema.graphql", Input: `directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
directive @goField(forceResolver: Boolean, name: String) on INPUT_FIELD_DEFINITION | FIELD_DEFINITION

type Query {
    invalidIdentifier: InvalidIdentifier
    collision: It
    mapInput(input: Changes): Boolean
    recursive(input: RecursiveInputSlice): Boolean
    nestedInputs(input: [[OuterInput]] = [[{inner: {id: 1}}]]): Boolean
    nestedOutputs: [[OuterObject]]
    modelMethods: ModelMethods
    user(id: Int!): User!
    nullableArg(arg: Int = 123): String
    inputSlice(arg: [String!]!): Boolean!
    shapeUnion: ShapeUnion!
    autobind: Autobind
    deprecatedField: String! @deprecated(reason: "test deprecated directive")
}

type Subscription {
    updated: String!
    initPayload: String!
}

type User {
    id: Int!
    friends: [User!]! @goField(forceResolver: true)
    created: Time!
    updated: Time
}

type Autobind {
    int: Int!
    int32: Int!
    int64: Int!

    idStr: ID!
    idInt: ID!
}

type ModelMethods {
    resolverField: Boolean!
    noContext: Boolean!
    withContext: Boolean!
}

type InvalidIdentifier {
    id: Int!
}

type It {
    id: ID!
}

input Changes @goModel(model:"map[string]interface{}") {
    a: Int
    b: Int
}

input RecursiveInputSlice {
    self: [RecursiveInputSlice!]
}

input InnerInput {
    id:Int!
}

input OuterInput {
    inner: InnerInput!
}

scalar ThirdParty @goModel(model:"testserver.ThirdParty")

type OuterObject {
    inner: InnerObject!
}

type InnerObject {
    id: Int!
}

type ForcedResolver {
    field: Circle @goField(forceResolver: true)
}

type EmbeddedPointer @goModel(model:"testserver.EmbeddedPointerModel") {
    ID: String
    Title: String
}

scalar UUID

enum Status {
    OK
    ERROR
}

scalar Time
`},
	&ast.Source{Name: "slices.graphql", Input: `extend type Query {
    slices: Slices
    scalarSlice: Bytes!
}

type Slices {
  test1: [String]
  test2: [String!]
  test3: [String]!
  test4: [String!]!
}

scalar Bytes
`},
	&ast.Source{Name: "typefallback.graphql", Input: `extend type Query {
    fallback(arg: FallbackToStringEncoding!): FallbackToStringEncoding!
}

enum FallbackToStringEncoding {
    A
    B
    C
}
`},
	&ast.Source{Name: "useptr.graphql", Input: `type A {
    id: ID!
}

type B {
    id: ID!
}

union TestUnion = A | B

extend type Query {
    optionalUnion: TestUnion
}
`},
	&ast.Source{Name: "validtypes.graphql", Input: `extend type Query {
    validType: ValidType
}

""" These things are all valid, but without care generate invalid go code """
type ValidType {
    differentCase: String!
    different_case: String! @goField(name:"DifferentCaseOld")
    validInputKeywords(input: ValidInput): Boolean!
    validArgs(
        break:       String!,
        default:     String!,
        func:        String!,
        interface:   String!,
        select:      String!,
        case:        String!,
        defer:       String!,
        go:          String!,
        map:         String!,
        struct:      String!,
        chan:        String!,
        else:        String!,
        goto:        String!,
        package:     String!,
        switch:      String!,
        const:       String!,
        fallthrough: String!,
        if:          String!,
        range:       String!,
        type:        String!,
        continue:    String!,
        for:         String!,
        import:      String!,
        return:      String!,
        var:         String!,
        _:           String!,
    ): Boolean!
}

input ValidInput {
    break:       String!
    default:     String!
    func:        String!
    interface:   String!
    select:      String!
    case:        String!
    defer:       String!
    go:          String!
    map:         String!
    struct:      String!
    chan:        String!
    else:        String!
    goto:        String!
    package:     String!
    switch:      String!
    const:       String!
    fallthrough: String!
    if:          String!
    range:       String!
    type:        String!
    continue:    String!
    for:         String!
    import:      String!
    return:      String!
    var:         String!
    _:           String! @goField(name: "Underscore")
}

# see https://github.com/99designs/gqlgen/issues/694
type Content_User {
  foo: String
}

type Content_Post {
  foo: String
}

union Content_Child = Content_User | Content_Post
`},
	&ast.Source{Name: "weird_type_cases.graphql", Input: `# regression test for https://github.com/99designs/gqlgen/issues/583

type asdfIt { id: ID! }
type iIt { id: ID! }
type AIt { id: ID! }
type XXIt { id: ID! }
type AbIt { id: ID! }
type XxIt { id: ID! }
`},
	&ast.Source{Name: "wrapped_type.graphql", Input: `# regression test for https://github.com/99designs/gqlgen/issues/721

extend type Query {
    wrappedStruct: WrappedStruct!
    wrappedScalar: WrappedScalar!
}

type WrappedStruct { name: String! }
scalar WrappedScalar
`},
)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) dir_length_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["min"]; ok {
		arg0, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["min"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["max"]; ok {
		arg1, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["max"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["message"]; ok {
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["message"] = arg2
	return args, nil
}

func (ec *executionContext) dir_logged_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNUUID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) dir_range_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["min"]; ok {
		arg0, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["min"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["max"]; ok {
		arg1, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["max"] = arg1
	return args, nil
}

func (ec *executionContext) field_Panics_argUnmarshal_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []MarshalPanic
	if tmp, ok := rawArgs["u"]; ok {
		arg0, err = ec.unmarshalNMarshalPanic2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["u"] = arg0
	return args, nil
}

func (ec *executionContext) field_Panics_fieldFuncMarshal_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []MarshalPanic
	if tmp, ok := rawArgs["u"]; ok {
		arg0, err = ec.unmarshalNMarshalPanic2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["u"] = arg0
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

func (ec *executionContext) field_Query_defaultScalar_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = ec.unmarshalNDefaultScalarImplementation2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_directiveArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["arg"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalNString2string(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalNInt2int(ctx, 1)
			if err != nil {
				return nil, err
			}
			max, err := ec.unmarshalOInt2ᚖint(ctx, 255)
			if err != nil {
				return nil, err
			}
			message, err := ec.unmarshalOString2ᚖstring(ctx, "invalid length")
			if err != nil {
				return nil, err
			}
			if ec.directives.Length == nil {
				return nil, errors.New("directive length is not implemented")
			}
			return ec.directives.Length(ctx, rawArgs, directive0, min, max, message)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(string); ok {
			arg0 = data
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_directiveFieldDef_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["ret"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["ret"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_directiveInputNullable_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *InputDirectives
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = ec.unmarshalOInputDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_directiveInputType_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 InnerInput
	if tmp, ok := rawArgs["arg"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) {
			return ec.unmarshalNInnerInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerInput(ctx, tmp)
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Custom == nil {
				return nil, errors.New("directive custom is not implemented")
			}
			return ec.directives.Custom(ctx, rawArgs, directive0)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(InnerInput); ok {
			arg0 = data
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be github.com/99designs/gqlgen/codegen/testserver.InnerInput`, tmp)
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_directiveInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 InputDirectives
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = ec.unmarshalNInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_directiveNullableArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["arg"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOInt2ᚖint(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalOInt2ᚖint(ctx, 0)
			if err != nil {
				return nil, err
			}
			if ec.directives.Range == nil {
				return nil, errors.New("directive range is not implemented")
			}
			return ec.directives.Range(ctx, rawArgs, directive0, min, nil)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*int); ok {
			arg0 = data
		} else if tmp == nil {
			arg0 = nil
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *int`, tmp)
		}
	}
	args["arg"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["arg2"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOInt2ᚖint(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalOInt2ᚖint(ctx, 0)
			if err != nil {
				return nil, err
			}
			if ec.directives.Range == nil {
				return nil, errors.New("directive range is not implemented")
			}
			return ec.directives.Range(ctx, rawArgs, directive0, min, nil)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*int); ok {
			arg1 = data
		} else if tmp == nil {
			arg1 = nil
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *int`, tmp)
		}
	}
	args["arg2"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["arg3"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOString2ᚖstring(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.ToNull == nil {
				return nil, errors.New("directive toNull is not implemented")
			}
			return ec.directives.ToNull(ctx, rawArgs, directive0)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*string); ok {
			arg2 = data
		} else if tmp == nil {
			arg2 = nil
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
		}
	}
	args["arg3"] = arg2
	return args, nil
}

func (ec *executionContext) field_Query_fallback_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 FallbackToStringEncoding
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = ec.unmarshalNFallbackToStringEncoding2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐFallbackToStringEncoding(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_inputSlice_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []string
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = ec.unmarshalNString2ᚕstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_mapInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 map[string]interface{}
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = ec.unmarshalOChanges2map(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_mapNestedStringInterface_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *NestedMapInput
	if tmp, ok := rawArgs["in"]; ok {
		arg0, err = ec.unmarshalONestedMapInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐNestedMapInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["in"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_mapStringInterface_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 map[string]interface{}
	if tmp, ok := rawArgs["in"]; ok {
		arg0, err = ec.unmarshalOMapStringInterfaceInput2map(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["in"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_nestedInputs_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 [][]*OuterInput
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = ec.unmarshalOOuterInput2ᚕᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_nullableArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_recursive_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *RecursiveInputSlice
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = ec.unmarshalORecursiveInputSlice2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx, tmp)
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
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Subscription_directiveArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["arg"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalNString2string(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalNInt2int(ctx, 1)
			if err != nil {
				return nil, err
			}
			max, err := ec.unmarshalOInt2ᚖint(ctx, 255)
			if err != nil {
				return nil, err
			}
			message, err := ec.unmarshalOString2ᚖstring(ctx, "invalid length")
			if err != nil {
				return nil, err
			}
			if ec.directives.Length == nil {
				return nil, errors.New("directive length is not implemented")
			}
			return ec.directives.Length(ctx, rawArgs, directive0, min, max, message)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(string); ok {
			arg0 = data
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (ec *executionContext) field_Subscription_directiveNullableArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["arg"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOInt2ᚖint(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalOInt2ᚖint(ctx, 0)
			if err != nil {
				return nil, err
			}
			if ec.directives.Range == nil {
				return nil, errors.New("directive range is not implemented")
			}
			return ec.directives.Range(ctx, rawArgs, directive0, min, nil)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*int); ok {
			arg0 = data
		} else if tmp == nil {
			arg0 = nil
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *int`, tmp)
		}
	}
	args["arg"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["arg2"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOInt2ᚖint(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalOInt2ᚖint(ctx, 0)
			if err != nil {
				return nil, err
			}
			if ec.directives.Range == nil {
				return nil, errors.New("directive range is not implemented")
			}
			return ec.directives.Range(ctx, rawArgs, directive0, min, nil)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*int); ok {
			arg1 = data
		} else if tmp == nil {
			arg1 = nil
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *int`, tmp)
		}
	}
	args["arg2"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["arg3"]; ok {
		directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOString2ᚖstring(ctx, tmp) }
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.ToNull == nil {
				return nil, errors.New("directive toNull is not implemented")
			}
			return ec.directives.ToNull(ctx, rawArgs, directive0)
		}

		tmp, err = directive1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*string); ok {
			arg2 = data
		} else if tmp == nil {
			arg2 = nil
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
		}
	}
	args["arg3"] = arg2
	return args, nil
}

func (ec *executionContext) field_ValidType_validArgs_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["break"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["break"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["default"]; ok {
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["default"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["func"]; ok {
		arg2, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["func"] = arg2
	var arg3 string
	if tmp, ok := rawArgs["interface"]; ok {
		arg3, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["interface"] = arg3
	var arg4 string
	if tmp, ok := rawArgs["select"]; ok {
		arg4, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["select"] = arg4
	var arg5 string
	if tmp, ok := rawArgs["case"]; ok {
		arg5, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["case"] = arg5
	var arg6 string
	if tmp, ok := rawArgs["defer"]; ok {
		arg6, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["defer"] = arg6
	var arg7 string
	if tmp, ok := rawArgs["go"]; ok {
		arg7, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["go"] = arg7
	var arg8 string
	if tmp, ok := rawArgs["map"]; ok {
		arg8, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["map"] = arg8
	var arg9 string
	if tmp, ok := rawArgs["struct"]; ok {
		arg9, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["struct"] = arg9
	var arg10 string
	if tmp, ok := rawArgs["chan"]; ok {
		arg10, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["chan"] = arg10
	var arg11 string
	if tmp, ok := rawArgs["else"]; ok {
		arg11, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["else"] = arg11
	var arg12 string
	if tmp, ok := rawArgs["goto"]; ok {
		arg12, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["goto"] = arg12
	var arg13 string
	if tmp, ok := rawArgs["package"]; ok {
		arg13, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["package"] = arg13
	var arg14 string
	if tmp, ok := rawArgs["switch"]; ok {
		arg14, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["switch"] = arg14
	var arg15 string
	if tmp, ok := rawArgs["const"]; ok {
		arg15, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["const"] = arg15
	var arg16 string
	if tmp, ok := rawArgs["fallthrough"]; ok {
		arg16, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["fallthrough"] = arg16
	var arg17 string
	if tmp, ok := rawArgs["if"]; ok {
		arg17, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["if"] = arg17
	var arg18 string
	if tmp, ok := rawArgs["range"]; ok {
		arg18, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["range"] = arg18
	var arg19 string
	if tmp, ok := rawArgs["type"]; ok {
		arg19, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["type"] = arg19
	var arg20 string
	if tmp, ok := rawArgs["continue"]; ok {
		arg20, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["continue"] = arg20
	var arg21 string
	if tmp, ok := rawArgs["for"]; ok {
		arg21, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["for"] = arg21
	var arg22 string
	if tmp, ok := rawArgs["import"]; ok {
		arg22, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["import"] = arg22
	var arg23 string
	if tmp, ok := rawArgs["return"]; ok {
		arg23, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["return"] = arg23
	var arg24 string
	if tmp, ok := rawArgs["var"]; ok {
		arg24, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["var"] = arg24
	var arg25 string
	if tmp, ok := rawArgs["_"]; ok {
		arg25, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["_"] = arg25
	return args, nil
}

func (ec *executionContext) field_ValidType_validInputKeywords_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *ValidInput
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = ec.unmarshalOValidInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
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

func (ec *executionContext) _fieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) interface{} {
	fc := graphql.GetFieldContext(ctx)
	for _, d := range fc.Field.Directives {
		switch d.Name {
		case "logged":
			rawArgs := d.ArgumentMap(ec.Variables)
			args, err := ec.dir_logged_args(ctx, rawArgs)
			if err != nil {
				ec.Error(ctx, err)
				return nil
			}
			n := next
			next = func(ctx context.Context) (interface{}, error) {
				if ec.directives.Logged == nil {
					return nil, errors.New("directive logged is not implemented")
				}
				return ec.directives.Logged(ctx, obj, n, args["id"].(string))
			}
		}
	}
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _A_id(ctx context.Context, field graphql.CollectedField, obj *A) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "A",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _AIt_id(ctx context.Context, field graphql.CollectedField, obj *AIt) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "AIt",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _AbIt_id(ctx context.Context, field graphql.CollectedField, obj *AbIt) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "AbIt",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Autobind_int(ctx context.Context, field graphql.CollectedField, obj *Autobind) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Autobind",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Int, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Autobind_int32(ctx context.Context, field graphql.CollectedField, obj *Autobind) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Autobind",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Int32, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int32)
	fc.Result = res
	return ec.marshalNInt2int32(ctx, field.Selections, res)
}

func (ec *executionContext) _Autobind_int64(ctx context.Context, field graphql.CollectedField, obj *Autobind) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Autobind",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Int64, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNInt2int64(ctx, field.Selections, res)
}

func (ec *executionContext) _Autobind_idStr(ctx context.Context, field graphql.CollectedField, obj *Autobind) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Autobind",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IdStr, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Autobind_idInt(ctx context.Context, field graphql.CollectedField, obj *Autobind) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Autobind",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IdInt, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) _B_id(ctx context.Context, field graphql.CollectedField, obj *B) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "B",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Circle_radius(ctx context.Context, field graphql.CollectedField, obj *Circle) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Circle",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Radius, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Circle_area(ctx context.Context, field graphql.CollectedField, obj *Circle) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Circle",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Area(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Content_Post_foo(ctx context.Context, field graphql.CollectedField, obj *ContentPost) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Content_Post",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Content_User_foo(ctx context.Context, field graphql.CollectedField, obj *ContentUser) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Content_User",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _EmbeddedDefaultScalar_value(ctx context.Context, field graphql.CollectedField, obj *EmbeddedDefaultScalar) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "EmbeddedDefaultScalar",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Value, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalODefaultScalarImplementation2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _EmbeddedPointer_ID(ctx context.Context, field graphql.CollectedField, obj *EmbeddedPointerModel) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "EmbeddedPointer",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _EmbeddedPointer_Title(ctx context.Context, field graphql.CollectedField, obj *EmbeddedPointerModel) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "EmbeddedPointer",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Title, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Error_id(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Error",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Error_errorOnNonRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Error",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorOnNonRequiredField()
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Error_errorOnRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Error",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorOnRequiredField()
	})

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

func (ec *executionContext) _Error_nilOnRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Error",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NilOnRequiredField(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalNString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Errors_a(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Errors",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().A(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) _Errors_b(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Errors",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().B(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) _Errors_c(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Errors",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().C(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) _Errors_d(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Errors",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().D(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) _Errors_e(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Errors",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().E(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) _ForcedResolver_field(ctx context.Context, field graphql.CollectedField, obj *ForcedResolver) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ForcedResolver",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.ForcedResolver().Field(rctx, obj)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Circle)
	fc.Result = res
	return ec.marshalOCircle2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐCircle(ctx, field.Selections, res)
}

func (ec *executionContext) _InnerObject_id(ctx context.Context, field graphql.CollectedField, obj *InnerObject) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "InnerObject",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _InvalidIdentifier_id(ctx context.Context, field graphql.CollectedField, obj *invalid_packagename.InvalidIdentifier) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "InvalidIdentifier",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _It_id(ctx context.Context, field graphql.CollectedField, obj *introspection1.It) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "It",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _LoopA_b(ctx context.Context, field graphql.CollectedField, obj *LoopA) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "LoopA",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.B, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*LoopB)
	fc.Result = res
	return ec.marshalNLoopB2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐLoopB(ctx, field.Selections, res)
}

func (ec *executionContext) _LoopB_a(ctx context.Context, field graphql.CollectedField, obj *LoopB) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "LoopB",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.A, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*LoopA)
	fc.Result = res
	return ec.marshalNLoopA2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐLoopA(ctx, field.Selections, res)
}

func (ec *executionContext) _Map_id(ctx context.Context, field graphql.CollectedField, obj *Map) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Map",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _MapStringInterfaceType_a(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MapStringInterfaceType",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["a"].(type) {
		case *string:
			return v, nil
		case string:
			return &v, nil
		case nil:
			return (*string)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "a")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _MapStringInterfaceType_b(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MapStringInterfaceType",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["b"].(type) {
		case *int:
			return v, nil
		case int:
			return &v, nil
		case nil:
			return (*int)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "b")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	fc.Result = res
	return ec.marshalOInt2ᚖint(ctx, field.Selections, res)
}

func (ec *executionContext) _ModelMethods_resolverField(ctx context.Context, field graphql.CollectedField, obj *ModelMethods) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ModelMethods",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.ModelMethods().ResolverField(rctx, obj)
	})

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

func (ec *executionContext) _ModelMethods_noContext(ctx context.Context, field graphql.CollectedField, obj *ModelMethods) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ModelMethods",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NoContext(), nil
	})

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

func (ec *executionContext) _ModelMethods_withContext(ctx context.Context, field graphql.CollectedField, obj *ModelMethods) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ModelMethods",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.WithContext(ctx), nil
	})

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

func (ec *executionContext) _ObjectDirectives_text(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectives) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ObjectDirectives",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.Text, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalNInt2int(ctx, 0)
			if err != nil {
				return nil, err
			}
			max, err := ec.unmarshalOInt2ᚖint(ctx, 7)
			if err != nil {
				return nil, err
			}
			message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
			if err != nil {
				return nil, err
			}
			if ec.directives.Length == nil {
				return nil, errors.New("directive length is not implemented")
			}
			return ec.directives.Length(ctx, obj, directive0, min, max, message)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
	})

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

func (ec *executionContext) _ObjectDirectives_nullableText(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectives) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ObjectDirectives",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.NullableText, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.ToNull == nil {
				return nil, errors.New("directive toNull is not implemented")
			}
			return ec.directives.ToNull(ctx, obj, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _ObjectDirectivesWithCustomGoModel_nullableText(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectivesWithCustomGoModel) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ObjectDirectivesWithCustomGoModel",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.NullableText, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.ToNull == nil {
				return nil, errors.New("directive toNull is not implemented")
			}
			return ec.directives.ToNull(ctx, obj, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _OuterObject_inner(ctx context.Context, field graphql.CollectedField, obj *OuterObject) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "OuterObject",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Inner, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*InnerObject)
	fc.Result = res
	return ec.marshalNInnerObject2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerObject(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_oneFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "OverlappingFields",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_twoFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "OverlappingFields",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_oldFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "OverlappingFields",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.OverlappingFields().OldFoo(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_newFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "OverlappingFields",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NewFoo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_new_foo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "OverlappingFields",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NewFoo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Panics_fieldScalarMarshal(ctx context.Context, field graphql.CollectedField, obj *Panics) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Panics",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Panics().FieldScalarMarshal(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]MarshalPanic)
	fc.Result = res
	return ec.marshalNMarshalPanic2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx, field.Selections, res)
}

func (ec *executionContext) _Panics_fieldFuncMarshal(ctx context.Context, field graphql.CollectedField, obj *Panics) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Panics",
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
	args, err := ec.field_Panics_fieldFuncMarshal_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.FieldFuncMarshal(ctx, args["u"].([]MarshalPanic)), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]MarshalPanic)
	fc.Result = res
	return ec.marshalNMarshalPanic2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx, field.Selections, res)
}

func (ec *executionContext) _Panics_argUnmarshal(ctx context.Context, field graphql.CollectedField, obj *Panics) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Panics",
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
	args, err := ec.field_Panics_argUnmarshal_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Panics().ArgUnmarshal(rctx, obj, args["u"].([]MarshalPanic))
	})

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

func (ec *executionContext) _Primitive_value(ctx context.Context, field graphql.CollectedField, obj *Primitive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Primitive",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Primitive().Value(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Primitive_squared(ctx context.Context, field graphql.CollectedField, obj *Primitive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Primitive",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Squared(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _PrimitiveString_value(ctx context.Context, field graphql.CollectedField, obj *PrimitiveString) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "PrimitiveString",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.PrimitiveString().Value(rctx, obj)
	})

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

func (ec *executionContext) _PrimitiveString_doubled(ctx context.Context, field graphql.CollectedField, obj *PrimitiveString) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "PrimitiveString",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Doubled(), nil
	})

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

func (ec *executionContext) _PrimitiveString_len(ctx context.Context, field graphql.CollectedField, obj *PrimitiveString) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "PrimitiveString",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.PrimitiveString().Len(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_invalidIdentifier(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().InvalidIdentifier(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*invalid_packagename.InvalidIdentifier)
	fc.Result = res
	return ec.marshalOInvalidIdentifier2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋinvalidᚑpackagenameᚐInvalidIdentifier(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_collision(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Collision(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection1.It)
	fc.Result = res
	return ec.marshalOIt2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋintrospectionᚐIt(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_mapInput(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_mapInput_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().MapInput(rctx, args["input"].(map[string]interface{}))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_recursive(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_recursive_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Recursive(rctx, args["input"].(*RecursiveInputSlice))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_nestedInputs(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_nestedInputs_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().NestedInputs(rctx, args["input"].([][]*OuterInput))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_nestedOutputs(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().NestedOutputs(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([][]*OuterObject)
	fc.Result = res
	return ec.marshalOOuterObject2ᚕᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_modelMethods(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ModelMethods(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ModelMethods)
	fc.Result = res
	return ec.marshalOModelMethods2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐModelMethods(ctx, field.Selections, res)
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().User(rctx, args["id"].(int))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*User)
	fc.Result = res
	return ec.marshalNUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_nullableArg(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_nullableArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().NullableArg(rctx, args["arg"].(*int))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_inputSlice(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_inputSlice_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().InputSlice(rctx, args["arg"].([]string))
	})

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

func (ec *executionContext) _Query_shapeUnion(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ShapeUnion(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(ShapeUnion)
	fc.Result = res
	return ec.marshalNShapeUnion2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShapeUnion(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_autobind(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Autobind(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Autobind)
	fc.Result = res
	return ec.marshalOAutobind2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐAutobind(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_deprecatedField(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DeprecatedField(rctx)
	})

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

func (ec *executionContext) _Query_overlapping(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Overlapping(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*OverlappingFields)
	fc.Result = res
	return ec.marshalOOverlappingFields2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOverlappingFields(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveArg(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_directiveArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveArg(rctx, args["arg"].(string))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveNullableArg(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_directiveNullableArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveNullableArg(rctx, args["arg"].(*int), args["arg2"].(*int), args["arg3"].(*string))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveInputNullable(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_directiveInputNullable_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveInputNullable(rctx, args["arg"].(*InputDirectives))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveInput(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_directiveInput_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveInput(rctx, args["arg"].(InputDirectives))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveInputType(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_directiveInputType_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveInputType(rctx, args["arg"].(InnerInput))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveObject(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveObject(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ObjectDirectives)
	fc.Result = res
	return ec.marshalOObjectDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐObjectDirectives(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveObjectWithCustomGoModel(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveObjectWithCustomGoModel(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ObjectDirectivesWithCustomGoModel)
	fc.Result = res
	return ec.marshalOObjectDirectivesWithCustomGoModel2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐObjectDirectivesWithCustomGoModel(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveFieldDef(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_directiveFieldDef_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Query().DirectiveFieldDef(rctx, args["ret"].(string))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalNInt2int(ctx, 1)
			if err != nil {
				return nil, err
			}
			message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
			if err != nil {
				return nil, err
			}
			if ec.directives.Length == nil {
				return nil, errors.New("directive length is not implemented")
			}
			return ec.directives.Length(ctx, nil, directive0, min, nil, message)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
	})

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

func (ec *executionContext) _Query_directiveField(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveField(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveDouble(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Query().DirectiveDouble(rctx)
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Directive1 == nil {
				return nil, errors.New("directive directive1 is not implemented")
			}
			return ec.directives.Directive1(ctx, nil, directive0)
		}
		directive2 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Directive2 == nil {
				return nil, errors.New("directive directive2 is not implemented")
			}
			return ec.directives.Directive2(ctx, nil, directive1)
		}

		tmp, err := directive2(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_directiveUnimplemented(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Query().DirectiveUnimplemented(rctx)
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Unimplemented == nil {
				return nil, errors.New("directive unimplemented is not implemented")
			}
			return ec.directives.Unimplemented(ctx, nil, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_shapes(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Shapes(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]Shape)
	fc.Result = res
	return ec.marshalOShape2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShape(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_noShape(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Query().NoShape(rctx)
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.MakeNil == nil {
				return nil, errors.New("directive makeNil is not implemented")
			}
			return ec.directives.MakeNil(ctx, nil, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(Shape); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be github.com/99designs/gqlgen/codegen/testserver.Shape`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Shape)
	fc.Result = res
	return ec.marshalOShape2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShape(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_mapStringInterface(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_mapStringInterface_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().MapStringInterface(rctx, args["in"].(map[string]interface{}))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(map[string]interface{})
	fc.Result = res
	return ec.marshalOMapStringInterfaceType2map(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_mapNestedStringInterface(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_mapNestedStringInterface_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().MapNestedStringInterface(rctx, args["in"].(*NestedMapInput))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(map[string]interface{})
	fc.Result = res
	return ec.marshalOMapStringInterfaceType2map(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_errorBubble(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ErrorBubble(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalOError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_errors(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Errors(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Errors)
	fc.Result = res
	return ec.marshalOErrors2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐErrors(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_valid(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Valid(rctx)
	})

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

func (ec *executionContext) _Query_panics(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Panics(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Panics)
	fc.Result = res
	return ec.marshalOPanics2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPanics(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_primitiveObject(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().PrimitiveObject(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]Primitive)
	fc.Result = res
	return ec.marshalNPrimitive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitive(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_primitiveStringObject(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().PrimitiveStringObject(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]PrimitiveString)
	fc.Result = res
	return ec.marshalNPrimitiveString2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitiveString(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_defaultScalar(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_defaultScalar_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DefaultScalar(rctx, args["arg"].(string))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNDefaultScalarImplementation2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_slices(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Slices(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Slices)
	fc.Result = res
	return ec.marshalOSlices2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐSlices(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_scalarSlice(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ScalarSlice(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]byte)
	fc.Result = res
	return ec.marshalNBytes2ᚕbyte(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_fallback(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	args, err := ec.field_Query_fallback_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Fallback(rctx, args["arg"].(FallbackToStringEncoding))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(FallbackToStringEncoding)
	fc.Result = res
	return ec.marshalNFallbackToStringEncoding2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐFallbackToStringEncoding(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_optionalUnion(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().OptionalUnion(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(TestUnion)
	fc.Result = res
	return ec.marshalOTestUnion2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐTestUnion(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_validType(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ValidType(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ValidType)
	fc.Result = res
	return ec.marshalOValidType2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_wrappedStruct(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().WrappedStruct(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*WrappedStruct)
	fc.Result = res
	return ec.marshalNWrappedStruct2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐWrappedStruct(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_wrappedScalar(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().WrappedScalar(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(WrappedScalar)
	fc.Result = res
	return ec.marshalNWrappedScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐWrappedScalar(ctx, field.Selections, res)
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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})

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
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_length(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Rectangle",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Length, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_width(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Rectangle",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Width, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_area(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Rectangle",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Area(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Slices_test1(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Slices",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test1, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*string)
	fc.Result = res
	return ec.marshalOString2ᚕᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Slices_test2(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Slices",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test2, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalOString2ᚕstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Slices_test3(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Slices",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test3, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*string)
	fc.Result = res
	return ec.marshalNString2ᚕᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Slices_test4(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Slices",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test4, nil
	})

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

func (ec *executionContext) _Subscription_updated(ctx context.Context, field graphql.CollectedField) (ret func() graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Subscription",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Subscription().Updated(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-resTmp.(<-chan string)
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			ec.marshalNString2string(ctx, field.Selections, res).MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _Subscription_initPayload(ctx context.Context, field graphql.CollectedField) (ret func() graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Subscription",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Subscription().InitPayload(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-resTmp.(<-chan string)
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			ec.marshalNString2string(ctx, field.Selections, res).MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _Subscription_directiveArg(ctx context.Context, field graphql.CollectedField) (ret func() graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Subscription",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Subscription_directiveArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Subscription().DirectiveArg(rctx, args["arg"].(string))
	})

	if resTmp == nil {
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-resTmp.(<-chan *string)
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			ec.marshalOString2ᚖstring(ctx, field.Selections, res).MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _Subscription_directiveNullableArg(ctx context.Context, field graphql.CollectedField) (ret func() graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Subscription",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Subscription_directiveNullableArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Subscription().DirectiveNullableArg(rctx, args["arg"].(*int), args["arg2"].(*int), args["arg3"].(*string))
	})

	if resTmp == nil {
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-resTmp.(<-chan *string)
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			ec.marshalOString2ᚖstring(ctx, field.Selections, res).MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _Subscription_directiveDouble(ctx context.Context, field graphql.CollectedField) (ret func() graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Subscription",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Subscription().DirectiveDouble(rctx)
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Directive1 == nil {
				return nil, errors.New("directive directive1 is not implemented")
			}
			return ec.directives.Directive1(ctx, nil, directive0)
		}
		directive2 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Directive2 == nil {
				return nil, errors.New("directive directive2 is not implemented")
			}
			return ec.directives.Directive2(ctx, nil, directive1)
		}

		tmp, err := directive2(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(<-chan *string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be <-chan *string`, tmp)
	})

	if resTmp == nil {
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-resTmp.(<-chan *string)
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			ec.marshalOString2ᚖstring(ctx, field.Selections, res).MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _Subscription_directiveUnimplemented(ctx context.Context, field graphql.CollectedField) (ret func() graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Subscription",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.Subscription().DirectiveUnimplemented(rctx)
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.Unimplemented == nil {
				return nil, errors.New("directive unimplemented is not implemented")
			}
			return ec.directives.Unimplemented(ctx, nil, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(<-chan *string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be <-chan *string`, tmp)
	})

	if resTmp == nil {
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-resTmp.(<-chan *string)
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			ec.marshalOString2ᚖstring(ctx, field.Selections, res).MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *User) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _User_friends(ctx context.Context, field graphql.CollectedField, obj *User) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.User().Friends(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*User)
	fc.Result = res
	return ec.marshalNUser2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) _User_created(ctx context.Context, field graphql.CollectedField, obj *User) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Created, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) _User_updated(ctx context.Context, field graphql.CollectedField, obj *User) (ret graphql.Marshaler) {
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Updated, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) _ValidType_differentCase(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ValidType",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DifferentCase, nil
	})

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

func (ec *executionContext) _ValidType_different_case(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ValidType",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DifferentCaseOld, nil
	})

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

func (ec *executionContext) _ValidType_validInputKeywords(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ValidType",
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
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_ValidType_validInputKeywords_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ValidInputKeywords, nil
	})

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

func (ec *executionContext) _ValidType_validArgs(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "ValidType",
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
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_ValidType_validArgs_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ValidArgs, nil
	})

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

func (ec *executionContext) _WrappedStruct_name(ctx context.Context, field graphql.CollectedField, obj *WrappedStruct) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "WrappedStruct",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

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

func (ec *executionContext) _XXIt_id(ctx context.Context, field graphql.CollectedField, obj *XXIt) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "XXIt",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _XxIt_id(ctx context.Context, field graphql.CollectedField, obj *XxIt) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "XxIt",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})

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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) _asdfIt_id(ctx context.Context, field graphql.CollectedField, obj *AsdfIt) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "asdfIt",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _iIt_id(ctx context.Context, field graphql.CollectedField, obj *IIt) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "iIt",
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
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputInnerDirectives(ctx context.Context, obj interface{}) (InnerDirectives, error) {
	var it InnerDirectives
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "message":
			var err error
			directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalNString2string(ctx, v) }
			directive1 := func(ctx context.Context) (interface{}, error) {
				min, err := ec.unmarshalNInt2int(ctx, 1)
				if err != nil {
					return nil, err
				}
				message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
				if err != nil {
					return nil, err
				}
				if ec.directives.Length == nil {
					return nil, errors.New("directive length is not implemented")
				}
				return ec.directives.Length(ctx, obj, directive0, min, nil, message)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, err
			}
			if data, ok := tmp.(string); ok {
				it.Message = data
			} else {
				return it, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputInnerInput(ctx context.Context, obj interface{}) (InnerInput, error) {
	var it InnerInput
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "id":
			var err error
			it.ID, err = ec.unmarshalNInt2int(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputInputDirectives(ctx context.Context, obj interface{}) (InputDirectives, error) {
	var it InputDirectives
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "text":
			var err error
			directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalNString2string(ctx, v) }
			directive1 := func(ctx context.Context) (interface{}, error) {
				min, err := ec.unmarshalNInt2int(ctx, 0)
				if err != nil {
					return nil, err
				}
				max, err := ec.unmarshalOInt2ᚖint(ctx, 7)
				if err != nil {
					return nil, err
				}
				message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
				if err != nil {
					return nil, err
				}
				if ec.directives.Length == nil {
					return nil, errors.New("directive length is not implemented")
				}
				return ec.directives.Length(ctx, obj, directive0, min, max, message)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, err
			}
			if data, ok := tmp.(string); ok {
				it.Text = data
			} else {
				return it, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
			}
		case "nullableText":
			var err error
			directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOString2ᚖstring(ctx, v) }
			directive1 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.ToNull == nil {
					return nil, errors.New("directive toNull is not implemented")
				}
				return ec.directives.ToNull(ctx, obj, directive0)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, err
			}
			if data, ok := tmp.(*string); ok {
				it.NullableText = data
			} else if tmp == nil {
				it.NullableText = nil
			} else {
				return it, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
			}
		case "inner":
			var err error
			it.Inner, err = ec.unmarshalNInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx, v)
			if err != nil {
				return it, err
			}
		case "innerNullable":
			var err error
			it.InnerNullable, err = ec.unmarshalOInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx, v)
			if err != nil {
				return it, err
			}
		case "thirdParty":
			var err error
			directive0 := func(ctx context.Context) (interface{}, error) {
				return ec.unmarshalOThirdParty2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx, v)
			}
			directive1 := func(ctx context.Context) (interface{}, error) {
				min, err := ec.unmarshalNInt2int(ctx, 0)
				if err != nil {
					return nil, err
				}
				max, err := ec.unmarshalOInt2ᚖint(ctx, 7)
				if err != nil {
					return nil, err
				}
				if ec.directives.Length == nil {
					return nil, errors.New("directive length is not implemented")
				}
				return ec.directives.Length(ctx, obj, directive0, min, max, nil)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, err
			}
			if data, ok := tmp.(*ThirdParty); ok {
				it.ThirdParty = data
			} else if tmp == nil {
				it.ThirdParty = nil
			} else {
				return it, fmt.Errorf(`unexpected type %T from directive, should be *github.com/99designs/gqlgen/codegen/testserver.ThirdParty`, tmp)
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputNestedMapInput(ctx context.Context, obj interface{}) (NestedMapInput, error) {
	var it NestedMapInput
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "map":
			var err error
			it.Map, err = ec.unmarshalOMapStringInterfaceInput2map(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputOuterInput(ctx context.Context, obj interface{}) (OuterInput, error) {
	var it OuterInput
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "inner":
			var err error
			it.Inner, err = ec.unmarshalNInnerInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerInput(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputRecursiveInputSlice(ctx context.Context, obj interface{}) (RecursiveInputSlice, error) {
	var it RecursiveInputSlice
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "self":
			var err error
			it.Self, err = ec.unmarshalORecursiveInputSlice2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputValidInput(ctx context.Context, obj interface{}) (ValidInput, error) {
	var it ValidInput
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "break":
			var err error
			it.Break, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "default":
			var err error
			it.Default, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "func":
			var err error
			it.Func, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "interface":
			var err error
			it.Interface, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "select":
			var err error
			it.Select, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "case":
			var err error
			it.Case, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "defer":
			var err error
			it.Defer, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "go":
			var err error
			it.Go, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "map":
			var err error
			it.Map, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "struct":
			var err error
			it.Struct, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "chan":
			var err error
			it.Chan, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "else":
			var err error
			it.Else, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "goto":
			var err error
			it.Goto, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "package":
			var err error
			it.Package, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "switch":
			var err error
			it.Switch, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "const":
			var err error
			it.Const, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "fallthrough":
			var err error
			it.Fallthrough, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "if":
			var err error
			it.If, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "range":
			var err error
			it.Range, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "type":
			var err error
			it.Type, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "continue":
			var err error
			it.Continue, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "for":
			var err error
			it.For, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "import":
			var err error
			it.Import, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "return":
			var err error
			it.Return, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "var":
			var err error
			it.Var, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "_":
			var err error
			it.Underscore, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Content_Child(ctx context.Context, sel ast.SelectionSet, obj ContentChild) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case ContentUser:
		return ec._Content_User(ctx, sel, &obj)
	case *ContentUser:
		return ec._Content_User(ctx, sel, obj)
	case ContentPost:
		return ec._Content_Post(ctx, sel, &obj)
	case *ContentPost:
		return ec._Content_Post(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _Shape(ctx context.Context, sel ast.SelectionSet, obj Shape) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _ShapeUnion(ctx context.Context, sel ast.SelectionSet, obj ShapeUnion) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _TestUnion(ctx context.Context, sel ast.SelectionSet, obj TestUnion) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case A:
		return ec._A(ctx, sel, &obj)
	case *A:
		return ec._A(ctx, sel, obj)
	case B:
		return ec._B(ctx, sel, &obj)
	case *B:
		return ec._B(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var aImplementors = []string{"A", "TestUnion"}

func (ec *executionContext) _A(ctx context.Context, sel ast.SelectionSet, obj *A) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, aImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("A")
		case "id":
			out.Values[i] = ec._A_id(ctx, field, obj)
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

var aItImplementors = []string{"AIt"}

func (ec *executionContext) _AIt(ctx context.Context, sel ast.SelectionSet, obj *AIt) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, aItImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("AIt")
		case "id":
			out.Values[i] = ec._AIt_id(ctx, field, obj)
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

var abItImplementors = []string{"AbIt"}

func (ec *executionContext) _AbIt(ctx context.Context, sel ast.SelectionSet, obj *AbIt) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, abItImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("AbIt")
		case "id":
			out.Values[i] = ec._AbIt_id(ctx, field, obj)
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

var autobindImplementors = []string{"Autobind"}

func (ec *executionContext) _Autobind(ctx context.Context, sel ast.SelectionSet, obj *Autobind) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, autobindImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Autobind")
		case "int":
			out.Values[i] = ec._Autobind_int(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "int32":
			out.Values[i] = ec._Autobind_int32(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "int64":
			out.Values[i] = ec._Autobind_int64(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "idStr":
			out.Values[i] = ec._Autobind_idStr(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "idInt":
			out.Values[i] = ec._Autobind_idInt(ctx, field, obj)
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

var bImplementors = []string{"B", "TestUnion"}

func (ec *executionContext) _B(ctx context.Context, sel ast.SelectionSet, obj *B) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, bImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("B")
		case "id":
			out.Values[i] = ec._B_id(ctx, field, obj)
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

var circleImplementors = []string{"Circle", "Shape", "ShapeUnion"}

func (ec *executionContext) _Circle(ctx context.Context, sel ast.SelectionSet, obj *Circle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, circleImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Circle")
		case "radius":
			out.Values[i] = ec._Circle_radius(ctx, field, obj)
		case "area":
			out.Values[i] = ec._Circle_area(ctx, field, obj)
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

var content_PostImplementors = []string{"Content_Post", "Content_Child"}

func (ec *executionContext) _Content_Post(ctx context.Context, sel ast.SelectionSet, obj *ContentPost) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, content_PostImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Content_Post")
		case "foo":
			out.Values[i] = ec._Content_Post_foo(ctx, field, obj)
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

var content_UserImplementors = []string{"Content_User", "Content_Child"}

func (ec *executionContext) _Content_User(ctx context.Context, sel ast.SelectionSet, obj *ContentUser) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, content_UserImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Content_User")
		case "foo":
			out.Values[i] = ec._Content_User_foo(ctx, field, obj)
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

var embeddedDefaultScalarImplementors = []string{"EmbeddedDefaultScalar"}

func (ec *executionContext) _EmbeddedDefaultScalar(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedDefaultScalar) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, embeddedDefaultScalarImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedDefaultScalar")
		case "value":
			out.Values[i] = ec._EmbeddedDefaultScalar_value(ctx, field, obj)
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

var embeddedPointerImplementors = []string{"EmbeddedPointer"}

func (ec *executionContext) _EmbeddedPointer(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedPointerModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, embeddedPointerImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedPointer")
		case "ID":
			out.Values[i] = ec._EmbeddedPointer_ID(ctx, field, obj)
		case "Title":
			out.Values[i] = ec._EmbeddedPointer_Title(ctx, field, obj)
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

var errorImplementors = []string{"Error"}

func (ec *executionContext) _Error(ctx context.Context, sel ast.SelectionSet, obj *Error) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, errorImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Error")
		case "id":
			out.Values[i] = ec._Error_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "errorOnNonRequiredField":
			out.Values[i] = ec._Error_errorOnNonRequiredField(ctx, field, obj)
		case "errorOnRequiredField":
			out.Values[i] = ec._Error_errorOnRequiredField(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "nilOnRequiredField":
			out.Values[i] = ec._Error_nilOnRequiredField(ctx, field, obj)
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

var errorsImplementors = []string{"Errors"}

func (ec *executionContext) _Errors(ctx context.Context, sel ast.SelectionSet, obj *Errors) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, errorsImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Errors")
		case "a":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_a(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "b":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_b(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "c":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_c(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "d":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_d(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "e":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_e(ctx, field, obj)
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

var forcedResolverImplementors = []string{"ForcedResolver"}

func (ec *executionContext) _ForcedResolver(ctx context.Context, sel ast.SelectionSet, obj *ForcedResolver) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, forcedResolverImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ForcedResolver")
		case "field":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._ForcedResolver_field(ctx, field, obj)
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

var innerObjectImplementors = []string{"InnerObject"}

func (ec *executionContext) _InnerObject(ctx context.Context, sel ast.SelectionSet, obj *InnerObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, innerObjectImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InnerObject")
		case "id":
			out.Values[i] = ec._InnerObject_id(ctx, field, obj)
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

var invalidIdentifierImplementors = []string{"InvalidIdentifier"}

func (ec *executionContext) _InvalidIdentifier(ctx context.Context, sel ast.SelectionSet, obj *invalid_packagename.InvalidIdentifier) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, invalidIdentifierImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InvalidIdentifier")
		case "id":
			out.Values[i] = ec._InvalidIdentifier_id(ctx, field, obj)
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

var itImplementors = []string{"It"}

func (ec *executionContext) _It(ctx context.Context, sel ast.SelectionSet, obj *introspection1.It) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, itImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("It")
		case "id":
			out.Values[i] = ec._It_id(ctx, field, obj)
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

var loopAImplementors = []string{"LoopA"}

func (ec *executionContext) _LoopA(ctx context.Context, sel ast.SelectionSet, obj *LoopA) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, loopAImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("LoopA")
		case "b":
			out.Values[i] = ec._LoopA_b(ctx, field, obj)
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

var loopBImplementors = []string{"LoopB"}

func (ec *executionContext) _LoopB(ctx context.Context, sel ast.SelectionSet, obj *LoopB) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, loopBImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("LoopB")
		case "a":
			out.Values[i] = ec._LoopB_a(ctx, field, obj)
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

var mapImplementors = []string{"Map"}

func (ec *executionContext) _Map(ctx context.Context, sel ast.SelectionSet, obj *Map) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mapImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Map")
		case "id":
			out.Values[i] = ec._Map_id(ctx, field, obj)
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

var mapStringInterfaceTypeImplementors = []string{"MapStringInterfaceType"}

func (ec *executionContext) _MapStringInterfaceType(ctx context.Context, sel ast.SelectionSet, obj map[string]interface{}) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mapStringInterfaceTypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MapStringInterfaceType")
		case "a":
			out.Values[i] = ec._MapStringInterfaceType_a(ctx, field, obj)
		case "b":
			out.Values[i] = ec._MapStringInterfaceType_b(ctx, field, obj)
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

var modelMethodsImplementors = []string{"ModelMethods"}

func (ec *executionContext) _ModelMethods(ctx context.Context, sel ast.SelectionSet, obj *ModelMethods) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, modelMethodsImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ModelMethods")
		case "resolverField":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._ModelMethods_resolverField(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "noContext":
			out.Values[i] = ec._ModelMethods_noContext(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "withContext":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._ModelMethods_withContext(ctx, field, obj)
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

var objectDirectivesImplementors = []string{"ObjectDirectives"}

func (ec *executionContext) _ObjectDirectives(ctx context.Context, sel ast.SelectionSet, obj *ObjectDirectives) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, objectDirectivesImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ObjectDirectives")
		case "text":
			out.Values[i] = ec._ObjectDirectives_text(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "nullableText":
			out.Values[i] = ec._ObjectDirectives_nullableText(ctx, field, obj)
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

var objectDirectivesWithCustomGoModelImplementors = []string{"ObjectDirectivesWithCustomGoModel"}

func (ec *executionContext) _ObjectDirectivesWithCustomGoModel(ctx context.Context, sel ast.SelectionSet, obj *ObjectDirectivesWithCustomGoModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, objectDirectivesWithCustomGoModelImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ObjectDirectivesWithCustomGoModel")
		case "nullableText":
			out.Values[i] = ec._ObjectDirectivesWithCustomGoModel_nullableText(ctx, field, obj)
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

var outerObjectImplementors = []string{"OuterObject"}

func (ec *executionContext) _OuterObject(ctx context.Context, sel ast.SelectionSet, obj *OuterObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, outerObjectImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OuterObject")
		case "inner":
			out.Values[i] = ec._OuterObject_inner(ctx, field, obj)
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

var overlappingFieldsImplementors = []string{"OverlappingFields"}

func (ec *executionContext) _OverlappingFields(ctx context.Context, sel ast.SelectionSet, obj *OverlappingFields) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, overlappingFieldsImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OverlappingFields")
		case "oneFoo":
			out.Values[i] = ec._OverlappingFields_oneFoo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "twoFoo":
			out.Values[i] = ec._OverlappingFields_twoFoo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "oldFoo":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._OverlappingFields_oldFoo(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "newFoo":
			out.Values[i] = ec._OverlappingFields_newFoo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "new_foo":
			out.Values[i] = ec._OverlappingFields_new_foo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
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

var panicsImplementors = []string{"Panics"}

func (ec *executionContext) _Panics(ctx context.Context, sel ast.SelectionSet, obj *Panics) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, panicsImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Panics")
		case "fieldScalarMarshal":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Panics_fieldScalarMarshal(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "fieldFuncMarshal":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Panics_fieldFuncMarshal(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "argUnmarshal":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Panics_argUnmarshal(ctx, field, obj)
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

var primitiveImplementors = []string{"Primitive"}

func (ec *executionContext) _Primitive(ctx context.Context, sel ast.SelectionSet, obj *Primitive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, primitiveImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Primitive")
		case "value":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Primitive_value(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "squared":
			out.Values[i] = ec._Primitive_squared(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
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

var primitiveStringImplementors = []string{"PrimitiveString"}

func (ec *executionContext) _PrimitiveString(ctx context.Context, sel ast.SelectionSet, obj *PrimitiveString) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, primitiveStringImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PrimitiveString")
		case "value":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._PrimitiveString_value(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "doubled":
			out.Values[i] = ec._PrimitiveString_doubled(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "len":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._PrimitiveString_len(ctx, field, obj)
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
		case "invalidIdentifier":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_invalidIdentifier(ctx, field)
				return res
			})
		case "collision":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_collision(ctx, field)
				return res
			})
		case "mapInput":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_mapInput(ctx, field)
				return res
			})
		case "recursive":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_recursive(ctx, field)
				return res
			})
		case "nestedInputs":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_nestedInputs(ctx, field)
				return res
			})
		case "nestedOutputs":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_nestedOutputs(ctx, field)
				return res
			})
		case "modelMethods":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_modelMethods(ctx, field)
				return res
			})
		case "user":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_user(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "nullableArg":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_nullableArg(ctx, field)
				return res
			})
		case "inputSlice":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_inputSlice(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "shapeUnion":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_shapeUnion(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "autobind":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_autobind(ctx, field)
				return res
			})
		case "deprecatedField":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_deprecatedField(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "overlapping":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_overlapping(ctx, field)
				return res
			})
		case "directiveArg":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveArg(ctx, field)
				return res
			})
		case "directiveNullableArg":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveNullableArg(ctx, field)
				return res
			})
		case "directiveInputNullable":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveInputNullable(ctx, field)
				return res
			})
		case "directiveInput":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveInput(ctx, field)
				return res
			})
		case "directiveInputType":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveInputType(ctx, field)
				return res
			})
		case "directiveObject":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveObject(ctx, field)
				return res
			})
		case "directiveObjectWithCustomGoModel":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveObjectWithCustomGoModel(ctx, field)
				return res
			})
		case "directiveFieldDef":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveFieldDef(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "directiveField":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveField(ctx, field)
				return res
			})
		case "directiveDouble":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveDouble(ctx, field)
				return res
			})
		case "directiveUnimplemented":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveUnimplemented(ctx, field)
				return res
			})
		case "shapes":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_shapes(ctx, field)
				return res
			})
		case "noShape":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_noShape(ctx, field)
				return res
			})
		case "mapStringInterface":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_mapStringInterface(ctx, field)
				return res
			})
		case "mapNestedStringInterface":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_mapNestedStringInterface(ctx, field)
				return res
			})
		case "errorBubble":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_errorBubble(ctx, field)
				return res
			})
		case "errors":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_errors(ctx, field)
				return res
			})
		case "valid":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_valid(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "panics":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_panics(ctx, field)
				return res
			})
		case "primitiveObject":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_primitiveObject(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "primitiveStringObject":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_primitiveStringObject(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "defaultScalar":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_defaultScalar(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "slices":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_slices(ctx, field)
				return res
			})
		case "scalarSlice":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_scalarSlice(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "fallback":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_fallback(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "optionalUnion":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_optionalUnion(ctx, field)
				return res
			})
		case "validType":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_validType(ctx, field)
				return res
			})
		case "wrappedStruct":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_wrappedStruct(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "wrappedScalar":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_wrappedScalar(ctx, field)
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

var rectangleImplementors = []string{"Rectangle", "Shape", "ShapeUnion"}

func (ec *executionContext) _Rectangle(ctx context.Context, sel ast.SelectionSet, obj *Rectangle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, rectangleImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Rectangle")
		case "length":
			out.Values[i] = ec._Rectangle_length(ctx, field, obj)
		case "width":
			out.Values[i] = ec._Rectangle_width(ctx, field, obj)
		case "area":
			out.Values[i] = ec._Rectangle_area(ctx, field, obj)
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

var slicesImplementors = []string{"Slices"}

func (ec *executionContext) _Slices(ctx context.Context, sel ast.SelectionSet, obj *Slices) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, slicesImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Slices")
		case "test1":
			out.Values[i] = ec._Slices_test1(ctx, field, obj)
		case "test2":
			out.Values[i] = ec._Slices_test2(ctx, field, obj)
		case "test3":
			out.Values[i] = ec._Slices_test3(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "test4":
			out.Values[i] = ec._Slices_test4(ctx, field, obj)
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

var subscriptionImplementors = []string{"Subscription"}

func (ec *executionContext) _Subscription(ctx context.Context, sel ast.SelectionSet) func() graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, subscriptionImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Subscription",
		Stats:  graphql.FieldStats{Started: graphql.Now()},
	})
	if len(fields) != 1 {
		ec.Errorf(ctx, "must subscribe to exactly one stream")
		return nil
	}

	switch fields[0].Name {
	case "updated":
		return ec._Subscription_updated(ctx, fields[0])
	case "initPayload":
		return ec._Subscription_initPayload(ctx, fields[0])
	case "directiveArg":
		return ec._Subscription_directiveArg(ctx, fields[0])
	case "directiveNullableArg":
		return ec._Subscription_directiveNullableArg(ctx, fields[0])
	case "directiveDouble":
		return ec._Subscription_directiveDouble(ctx, fields[0])
	case "directiveUnimplemented":
		return ec._Subscription_directiveUnimplemented(ctx, fields[0])
	default:
		panic("unknown field " + strconv.Quote(fields[0].Name))
	}
}

var userImplementors = []string{"User"}

func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *User) graphql.Marshaler {
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
		case "friends":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._User_friends(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "created":
			out.Values[i] = ec._User_created(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "updated":
			out.Values[i] = ec._User_updated(ctx, field, obj)
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

var validTypeImplementors = []string{"ValidType"}

func (ec *executionContext) _ValidType(ctx context.Context, sel ast.SelectionSet, obj *ValidType) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, validTypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ValidType")
		case "differentCase":
			out.Values[i] = ec._ValidType_differentCase(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "different_case":
			out.Values[i] = ec._ValidType_different_case(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "validInputKeywords":
			out.Values[i] = ec._ValidType_validInputKeywords(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "validArgs":
			out.Values[i] = ec._ValidType_validArgs(ctx, field, obj)
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

var wrappedStructImplementors = []string{"WrappedStruct"}

func (ec *executionContext) _WrappedStruct(ctx context.Context, sel ast.SelectionSet, obj *WrappedStruct) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, wrappedStructImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("WrappedStruct")
		case "name":
			out.Values[i] = ec._WrappedStruct_name(ctx, field, obj)
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

var xXItImplementors = []string{"XXIt"}

func (ec *executionContext) _XXIt(ctx context.Context, sel ast.SelectionSet, obj *XXIt) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, xXItImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("XXIt")
		case "id":
			out.Values[i] = ec._XXIt_id(ctx, field, obj)
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

var xxItImplementors = []string{"XxIt"}

func (ec *executionContext) _XxIt(ctx context.Context, sel ast.SelectionSet, obj *XxIt) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, xxItImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("XxIt")
		case "id":
			out.Values[i] = ec._XxIt_id(ctx, field, obj)
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

var asdfItImplementors = []string{"asdfIt"}

func (ec *executionContext) _asdfIt(ctx context.Context, sel ast.SelectionSet, obj *AsdfIt) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, asdfItImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("asdfIt")
		case "id":
			out.Values[i] = ec._asdfIt_id(ctx, field, obj)
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

var iItImplementors = []string{"iIt"}

func (ec *executionContext) _iIt(ctx context.Context, sel ast.SelectionSet, obj *IIt) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, iItImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("iIt")
		case "id":
			out.Values[i] = ec._iIt_id(ctx, field, obj)
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

func (ec *executionContext) unmarshalNBytes2ᚕbyte(ctx context.Context, v interface{}) ([]byte, error) {
	return UnmarshalBytes(v)
}

func (ec *executionContext) marshalNBytes2ᚕbyte(ctx context.Context, sel ast.SelectionSet, v []byte) graphql.Marshaler {
	res := MarshalBytes(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNDefaultScalarImplementation2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalNDefaultScalarImplementation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNError2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx context.Context, sel ast.SelectionSet, v Error) graphql.Marshaler {
	return ec._Error(ctx, sel, &v)
}

func (ec *executionContext) marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx context.Context, sel ast.SelectionSet, v *Error) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Error(ctx, sel, v)
}

func (ec *executionContext) unmarshalNFallbackToStringEncoding2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐFallbackToStringEncoding(ctx context.Context, v interface{}) (FallbackToStringEncoding, error) {
	tmp, err := graphql.UnmarshalString(v)
	return FallbackToStringEncoding(tmp), err
}

func (ec *executionContext) marshalNFallbackToStringEncoding2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐFallbackToStringEncoding(ctx context.Context, sel ast.SelectionSet, v FallbackToStringEncoding) graphql.Marshaler {
	res := graphql.MarshalString(string(v))
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNID2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalIntID(v)
}

func (ec *executionContext) marshalNID2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	res := graphql.MarshalIntID(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNID2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalID(v)
}

func (ec *executionContext) marshalNID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalID(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNInnerDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx context.Context, v interface{}) (InnerDirectives, error) {
	return ec.unmarshalInputInnerDirectives(ctx, v)
}

func (ec *executionContext) unmarshalNInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx context.Context, v interface{}) (*InnerDirectives, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalNInnerDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx, v)
	return &res, err
}

func (ec *executionContext) unmarshalNInnerInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerInput(ctx context.Context, v interface{}) (InnerInput, error) {
	return ec.unmarshalInputInnerInput(ctx, v)
}

func (ec *executionContext) unmarshalNInnerInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerInput(ctx context.Context, v interface{}) (*InnerInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalNInnerInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerInput(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalNInnerObject2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerObject(ctx context.Context, sel ast.SelectionSet, v InnerObject) graphql.Marshaler {
	return ec._InnerObject(ctx, sel, &v)
}

func (ec *executionContext) marshalNInnerObject2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerObject(ctx context.Context, sel ast.SelectionSet, v *InnerObject) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._InnerObject(ctx, sel, v)
}

func (ec *executionContext) unmarshalNInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(ctx context.Context, v interface{}) (InputDirectives, error) {
	return ec.unmarshalInputInputDirectives(ctx, v)
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

func (ec *executionContext) unmarshalNInt2int32(ctx context.Context, v interface{}) (int32, error) {
	return graphql.UnmarshalInt32(v)
}

func (ec *executionContext) marshalNInt2int32(ctx context.Context, sel ast.SelectionSet, v int32) graphql.Marshaler {
	res := graphql.MarshalInt32(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNInt2int64(ctx context.Context, v interface{}) (int64, error) {
	return graphql.UnmarshalInt64(v)
}

func (ec *executionContext) marshalNInt2int64(ctx context.Context, sel ast.SelectionSet, v int64) graphql.Marshaler {
	res := graphql.MarshalInt64(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNLoopA2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐLoopA(ctx context.Context, sel ast.SelectionSet, v LoopA) graphql.Marshaler {
	return ec._LoopA(ctx, sel, &v)
}

func (ec *executionContext) marshalNLoopA2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐLoopA(ctx context.Context, sel ast.SelectionSet, v *LoopA) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._LoopA(ctx, sel, v)
}

func (ec *executionContext) marshalNLoopB2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐLoopB(ctx context.Context, sel ast.SelectionSet, v LoopB) graphql.Marshaler {
	return ec._LoopB(ctx, sel, &v)
}

func (ec *executionContext) marshalNLoopB2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐLoopB(ctx context.Context, sel ast.SelectionSet, v *LoopB) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._LoopB(ctx, sel, v)
}

func (ec *executionContext) unmarshalNMarshalPanic2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx context.Context, v interface{}) (MarshalPanic, error) {
	var res MarshalPanic
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalNMarshalPanic2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx context.Context, sel ast.SelectionSet, v MarshalPanic) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNMarshalPanic2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx context.Context, v interface{}) ([]MarshalPanic, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]MarshalPanic, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNMarshalPanic2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNMarshalPanic2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx context.Context, sel ast.SelectionSet, v []MarshalPanic) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNMarshalPanic2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐMarshalPanic(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) marshalNPrimitive2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitive(ctx context.Context, sel ast.SelectionSet, v Primitive) graphql.Marshaler {
	return ec._Primitive(ctx, sel, &v)
}

func (ec *executionContext) marshalNPrimitive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitive(ctx context.Context, sel ast.SelectionSet, v []Primitive) graphql.Marshaler {
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
			ret[i] = ec.marshalNPrimitive2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitive(ctx, sel, v[i])
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

func (ec *executionContext) marshalNPrimitiveString2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitiveString(ctx context.Context, sel ast.SelectionSet, v PrimitiveString) graphql.Marshaler {
	return ec._PrimitiveString(ctx, sel, &v)
}

func (ec *executionContext) marshalNPrimitiveString2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitiveString(ctx context.Context, sel ast.SelectionSet, v []PrimitiveString) graphql.Marshaler {
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
			ret[i] = ec.marshalNPrimitiveString2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPrimitiveString(ctx, sel, v[i])
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

func (ec *executionContext) unmarshalNRecursiveInputSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx context.Context, v interface{}) (RecursiveInputSlice, error) {
	return ec.unmarshalInputRecursiveInputSlice(ctx, v)
}

func (ec *executionContext) marshalNShapeUnion2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShapeUnion(ctx context.Context, sel ast.SelectionSet, v ShapeUnion) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._ShapeUnion(ctx, sel, v)
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

func (ec *executionContext) unmarshalNString2ᚕᚖstring(ctx context.Context, v interface{}) ([]*string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]*string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalOString2ᚖstring(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNString2ᚕᚖstring(ctx context.Context, sel ast.SelectionSet, v []*string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalOString2ᚖstring(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) unmarshalNString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalNString2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalNString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec.marshalNString2string(ctx, sel, *v)
}

func (ec *executionContext) unmarshalNTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}

func (ec *executionContext) marshalNTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := graphql.MarshalTime(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNUUID2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalNUUID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNUser2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐUser(ctx context.Context, sel ast.SelectionSet, v User) graphql.Marshaler {
	return ec._User(ctx, sel, &v)
}

func (ec *executionContext) marshalNUser2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐUser(ctx context.Context, sel ast.SelectionSet, v []*User) graphql.Marshaler {
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
			ret[i] = ec.marshalNUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐUser(ctx, sel, v[i])
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

func (ec *executionContext) marshalNUser2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐUser(ctx context.Context, sel ast.SelectionSet, v *User) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
}

func (ec *executionContext) unmarshalNWrappedScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐWrappedScalar(ctx context.Context, v interface{}) (WrappedScalar, error) {
	tmp, err := graphql.UnmarshalString(v)
	return WrappedScalar(tmp), err
}

func (ec *executionContext) marshalNWrappedScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐWrappedScalar(ctx context.Context, sel ast.SelectionSet, v WrappedScalar) graphql.Marshaler {
	res := graphql.MarshalString(string(v))
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNWrappedStruct2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐWrappedStruct(ctx context.Context, sel ast.SelectionSet, v WrappedStruct) graphql.Marshaler {
	return ec._WrappedStruct(ctx, sel, &v)
}

func (ec *executionContext) marshalNWrappedStruct2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐWrappedStruct(ctx context.Context, sel ast.SelectionSet, v *WrappedStruct) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._WrappedStruct(ctx, sel, v)
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

func (ec *executionContext) marshalOAutobind2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐAutobind(ctx context.Context, sel ast.SelectionSet, v Autobind) graphql.Marshaler {
	return ec._Autobind(ctx, sel, &v)
}

func (ec *executionContext) marshalOAutobind2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐAutobind(ctx context.Context, sel ast.SelectionSet, v *Autobind) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Autobind(ctx, sel, v)
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

func (ec *executionContext) unmarshalOChanges2map(ctx context.Context, v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, nil
	}
	return v.(map[string]interface{}), nil
}

func (ec *executionContext) marshalOCircle2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐCircle(ctx context.Context, sel ast.SelectionSet, v Circle) graphql.Marshaler {
	return ec._Circle(ctx, sel, &v)
}

func (ec *executionContext) marshalOCircle2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐCircle(ctx context.Context, sel ast.SelectionSet, v *Circle) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Circle(ctx, sel, v)
}

func (ec *executionContext) unmarshalODefaultScalarImplementation2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalODefaultScalarImplementation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalODefaultScalarImplementation2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalODefaultScalarImplementation2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalODefaultScalarImplementation2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalODefaultScalarImplementation2string(ctx, sel, *v)
}

func (ec *executionContext) marshalOError2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx context.Context, sel ast.SelectionSet, v Error) graphql.Marshaler {
	return ec._Error(ctx, sel, &v)
}

func (ec *executionContext) marshalOError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐError(ctx context.Context, sel ast.SelectionSet, v *Error) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Error(ctx, sel, v)
}

func (ec *executionContext) marshalOErrors2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐErrors(ctx context.Context, sel ast.SelectionSet, v Errors) graphql.Marshaler {
	return ec._Errors(ctx, sel, &v)
}

func (ec *executionContext) marshalOErrors2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐErrors(ctx context.Context, sel ast.SelectionSet, v *Errors) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Errors(ctx, sel, v)
}

func (ec *executionContext) unmarshalOFloat2float64(ctx context.Context, v interface{}) (float64, error) {
	return graphql.UnmarshalFloat(v)
}

func (ec *executionContext) marshalOFloat2float64(ctx context.Context, sel ast.SelectionSet, v float64) graphql.Marshaler {
	return graphql.MarshalFloat(v)
}

func (ec *executionContext) unmarshalOInnerDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx context.Context, v interface{}) (InnerDirectives, error) {
	return ec.unmarshalInputInnerDirectives(ctx, v)
}

func (ec *executionContext) unmarshalOInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx context.Context, v interface{}) (*InnerDirectives, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOInnerDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInnerDirectives(ctx, v)
	return &res, err
}

func (ec *executionContext) unmarshalOInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(ctx context.Context, v interface{}) (InputDirectives, error) {
	return ec.unmarshalInputInputDirectives(ctx, v)
}

func (ec *executionContext) unmarshalOInputDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(ctx context.Context, v interface{}) (*InputDirectives, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(ctx, v)
	return &res, err
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

func (ec *executionContext) marshalOInvalidIdentifier2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋinvalidᚑpackagenameᚐInvalidIdentifier(ctx context.Context, sel ast.SelectionSet, v invalid_packagename.InvalidIdentifier) graphql.Marshaler {
	return ec._InvalidIdentifier(ctx, sel, &v)
}

func (ec *executionContext) marshalOInvalidIdentifier2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋinvalidᚑpackagenameᚐInvalidIdentifier(ctx context.Context, sel ast.SelectionSet, v *invalid_packagename.InvalidIdentifier) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._InvalidIdentifier(ctx, sel, v)
}

func (ec *executionContext) marshalOIt2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋintrospectionᚐIt(ctx context.Context, sel ast.SelectionSet, v introspection1.It) graphql.Marshaler {
	return ec._It(ctx, sel, &v)
}

func (ec *executionContext) marshalOIt2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋintrospectionᚐIt(ctx context.Context, sel ast.SelectionSet, v *introspection1.It) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._It(ctx, sel, v)
}

func (ec *executionContext) unmarshalOMapStringInterfaceInput2map(ctx context.Context, v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, nil
	}
	return v.(map[string]interface{}), nil
}

func (ec *executionContext) marshalOMapStringInterfaceType2map(ctx context.Context, sel ast.SelectionSet, v map[string]interface{}) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._MapStringInterfaceType(ctx, sel, v)
}

func (ec *executionContext) marshalOModelMethods2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐModelMethods(ctx context.Context, sel ast.SelectionSet, v ModelMethods) graphql.Marshaler {
	return ec._ModelMethods(ctx, sel, &v)
}

func (ec *executionContext) marshalOModelMethods2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐModelMethods(ctx context.Context, sel ast.SelectionSet, v *ModelMethods) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ModelMethods(ctx, sel, v)
}

func (ec *executionContext) unmarshalONestedMapInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐNestedMapInput(ctx context.Context, v interface{}) (NestedMapInput, error) {
	return ec.unmarshalInputNestedMapInput(ctx, v)
}

func (ec *executionContext) unmarshalONestedMapInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐNestedMapInput(ctx context.Context, v interface{}) (*NestedMapInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalONestedMapInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐNestedMapInput(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOObjectDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐObjectDirectives(ctx context.Context, sel ast.SelectionSet, v ObjectDirectives) graphql.Marshaler {
	return ec._ObjectDirectives(ctx, sel, &v)
}

func (ec *executionContext) marshalOObjectDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐObjectDirectives(ctx context.Context, sel ast.SelectionSet, v *ObjectDirectives) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ObjectDirectives(ctx, sel, v)
}

func (ec *executionContext) marshalOObjectDirectivesWithCustomGoModel2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐObjectDirectivesWithCustomGoModel(ctx context.Context, sel ast.SelectionSet, v ObjectDirectivesWithCustomGoModel) graphql.Marshaler {
	return ec._ObjectDirectivesWithCustomGoModel(ctx, sel, &v)
}

func (ec *executionContext) marshalOObjectDirectivesWithCustomGoModel2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐObjectDirectivesWithCustomGoModel(ctx context.Context, sel ast.SelectionSet, v *ObjectDirectivesWithCustomGoModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ObjectDirectivesWithCustomGoModel(ctx, sel, v)
}

func (ec *executionContext) unmarshalOOuterInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx context.Context, v interface{}) (OuterInput, error) {
	return ec.unmarshalInputOuterInput(ctx, v)
}

func (ec *executionContext) unmarshalOOuterInput2ᚕᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx context.Context, v interface{}) ([][]*OuterInput, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([][]*OuterInput, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalOOuterInput2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalOOuterInput2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx context.Context, v interface{}) ([]*OuterInput, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]*OuterInput, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalOOuterInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalOOuterInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx context.Context, v interface{}) (*OuterInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOOuterInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOOuterObject2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx context.Context, sel ast.SelectionSet, v OuterObject) graphql.Marshaler {
	return ec._OuterObject(ctx, sel, &v)
}

func (ec *executionContext) marshalOOuterObject2ᚕᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx context.Context, sel ast.SelectionSet, v [][]*OuterObject) graphql.Marshaler {
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
			ret[i] = ec.marshalOOuterObject2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx, sel, v[i])
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

func (ec *executionContext) marshalOOuterObject2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx context.Context, sel ast.SelectionSet, v []*OuterObject) graphql.Marshaler {
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
			ret[i] = ec.marshalOOuterObject2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx, sel, v[i])
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

func (ec *executionContext) marshalOOuterObject2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterObject(ctx context.Context, sel ast.SelectionSet, v *OuterObject) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._OuterObject(ctx, sel, v)
}

func (ec *executionContext) marshalOOverlappingFields2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOverlappingFields(ctx context.Context, sel ast.SelectionSet, v OverlappingFields) graphql.Marshaler {
	return ec._OverlappingFields(ctx, sel, &v)
}

func (ec *executionContext) marshalOOverlappingFields2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOverlappingFields(ctx context.Context, sel ast.SelectionSet, v *OverlappingFields) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._OverlappingFields(ctx, sel, v)
}

func (ec *executionContext) marshalOPanics2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPanics(ctx context.Context, sel ast.SelectionSet, v Panics) graphql.Marshaler {
	return ec._Panics(ctx, sel, &v)
}

func (ec *executionContext) marshalOPanics2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐPanics(ctx context.Context, sel ast.SelectionSet, v *Panics) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Panics(ctx, sel, v)
}

func (ec *executionContext) unmarshalORecursiveInputSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx context.Context, v interface{}) (RecursiveInputSlice, error) {
	return ec.unmarshalInputRecursiveInputSlice(ctx, v)
}

func (ec *executionContext) unmarshalORecursiveInputSlice2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx context.Context, v interface{}) ([]RecursiveInputSlice, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]RecursiveInputSlice, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNRecursiveInputSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalORecursiveInputSlice2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx context.Context, v interface{}) (*RecursiveInputSlice, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalORecursiveInputSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOShape2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShape(ctx context.Context, sel ast.SelectionSet, v Shape) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Shape(ctx, sel, v)
}

func (ec *executionContext) marshalOShape2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShape(ctx context.Context, sel ast.SelectionSet, v []Shape) graphql.Marshaler {
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
			ret[i] = ec.marshalOShape2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐShape(ctx, sel, v[i])
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

func (ec *executionContext) marshalOSlices2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐSlices(ctx context.Context, sel ast.SelectionSet, v Slices) graphql.Marshaler {
	return ec._Slices(ctx, sel, &v)
}

func (ec *executionContext) marshalOSlices2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐSlices(ctx context.Context, sel ast.SelectionSet, v *Slices) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Slices(ctx, sel, v)
}

func (ec *executionContext) unmarshalOString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalOString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOString2ᚕstring(ctx context.Context, v interface{}) ([]string, error) {
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

func (ec *executionContext) marshalOString2ᚕstring(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNString2string(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) unmarshalOString2ᚕᚖstring(ctx context.Context, v interface{}) ([]*string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]*string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalOString2ᚖstring(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOString2ᚕᚖstring(ctx context.Context, sel ast.SelectionSet, v []*string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalOString2ᚖstring(ctx, sel, v[i])
	}

	return ret
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

func (ec *executionContext) marshalOTestUnion2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐTestUnion(ctx context.Context, sel ast.SelectionSet, v TestUnion) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._TestUnion(ctx, sel, v)
}

func (ec *executionContext) unmarshalOThirdParty2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx context.Context, v interface{}) (ThirdParty, error) {
	return UnmarshalThirdParty(v)
}

func (ec *executionContext) marshalOThirdParty2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx context.Context, sel ast.SelectionSet, v ThirdParty) graphql.Marshaler {
	return MarshalThirdParty(v)
}

func (ec *executionContext) unmarshalOThirdParty2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx context.Context, v interface{}) (*ThirdParty, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOThirdParty2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOThirdParty2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx context.Context, sel ast.SelectionSet, v *ThirdParty) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOThirdParty2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐThirdParty(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}

func (ec *executionContext) marshalOTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	return graphql.MarshalTime(v)
}

func (ec *executionContext) unmarshalOTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOTime2timeᚐTime(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOTime2timeᚐTime(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOValidInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidInput(ctx context.Context, v interface{}) (ValidInput, error) {
	return ec.unmarshalInputValidInput(ctx, v)
}

func (ec *executionContext) unmarshalOValidInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidInput(ctx context.Context, v interface{}) (*ValidInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOValidInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidInput(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOValidType2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidType(ctx context.Context, sel ast.SelectionSet, v ValidType) graphql.Marshaler {
	return ec._ValidType(ctx, sel, &v)
}

func (ec *executionContext) marshalOValidType2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐValidType(ctx context.Context, sel ast.SelectionSet, v *ValidType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ValidType(ctx, sel, v)
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
