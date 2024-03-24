package configurationmodel

type ValueType int

const (
	IntSetType ValueType = iota
	IntRangeType
	FinalInt
	StringSetType
)

type ValueModel struct {
	valueType        ValueType
	intValues        []int
	stringValues     []string
	min, max         int
	minOpen, maxOpen bool
	finalValue       int
}

func NewIntSetModel(values []int) ValueModel {
	return ValueModel{
		valueType: IntSetType,
		intValues: values,
	}
}

func NewIntRangeModel(min int, minOpen bool, max int, maxOpen bool) ValueModel {
	return ValueModel{
		valueType: IntRangeType,
		min:       min,
		minOpen:   minOpen,
		max:       max,
		maxOpen:   maxOpen,
	}
}

func NewFinalIntModel(value int) ValueModel {
	return ValueModel{
		valueType:  FinalInt,
		finalValue: value,
	}
}

func NewStringSetModel(values []string) ValueModel {
	return ValueModel{
		valueType:    StringSetType,
		stringValues: values,
	}
}

type ParameterModel struct {
	id    int
	name  string
	value ValueModel
}

func (pModel ParameterModel) Id() int {
	return pModel.id
}

type ConstraintType int

const (
	SetValueIfFinal ConstraintType = iota
	SetValueIfValue
	ExcludeValueIfValue
)

type ConstraintModel struct {
	constraintType        ConstraintType
	srcId, targetId       int
	srcValue, targetValue ValueModel
}

func NewSetValueIfFinalConstraintModel(srcId, targetId int, targetValue ValueModel) ConstraintModel {
	return ConstraintModel{
		constraintType: SetValueIfFinal,
		srcId:          srcId,
		targetId:       targetId,
		targetValue:    targetValue,
	}
}

func NewSetValueIfValueConstraintModel(srcId int, srcValue ValueModel, targetId int, targetValue ValueModel) ConstraintModel {
	return ConstraintModel{
		constraintType: SetValueIfValue,
		srcId:          srcId,
		targetId:       targetId,
		srcValue:       srcValue,
		targetValue:    targetValue,
	}
}

func NewExcludeValueIfValueConstraintModel(srcId int, srcValue ValueModel, targetId int, targetValue ValueModel) ConstraintModel {
	return ConstraintModel{
		constraintType: ExcludeValueIfValue,
		srcId:          srcId,
		targetId:       targetId,
		srcValue:       srcValue,
		targetValue:    targetValue,
	}
}

func (cModel ConstraintModel) ConstraintType() ConstraintType {
	return cModel.constraintType
}

func (cModel ConstraintModel) SrcId() int {
	return cModel.srcId
}

func (cModel ConstraintModel) SrcValue() ValueModel {
	return cModel.srcValue
}

func (cModel ConstraintModel) TargetId() int {
	return cModel.targetId
}

func (cModel ConstraintModel) TargetValue() ValueModel {
	return cModel.targetValue
}

type Model struct {
	nextParameterId int
	parameters      []ParameterModel
	constraints     []ConstraintModel
}

func (model *Model) AddParameter(name string, value ValueModel) ParameterModel {
	model.nextParameterId++
	newParameter := ParameterModel{
		id:    model.nextParameterId,
		name:  name,
		value: value,
	}
	model.parameters = append(model.parameters, newParameter)
	return newParameter
}

func (model *Model) AddConstraint(constraint ConstraintModel) {
	model.constraints = append(model.constraints, constraint)
}

func (model Model) Parameters() []ParameterModel {
	return model.parameters
}

func (model Model) Constraints() []ConstraintModel {
	return model.constraints
}
