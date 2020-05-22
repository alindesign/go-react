package react

import (
	"github.com/iancoleman/strcase"
	"github.com/spf13/cast"
	"strings"
)

type Style struct {
	AlignContent                   interface{}
	AlignItems                     interface{}
	AlignSelf                      interface{}
	AlignmentBaseline              interface{}
	All                            interface{}
	Animation                      interface{}
	AnimationDelay                 interface{}
	AnimationDirection             interface{}
	AnimationDuration              interface{}
	AnimationFillMode              interface{}
	AnimationIterationCount        interface{}
	AnimationName                  interface{}
	AnimationPlayState             interface{}
	AnimationTimingFunction        interface{}
	BackdropFilter                 interface{}
	BackfaceVisibility             interface{}
	Background                     interface{}
	BackgroundAttachment           interface{}
	BackgroundBlendMode            interface{}
	BackgroundClip                 interface{}
	BackgroundColor                interface{}
	BackgroundImage                interface{}
	BackgroundOrigin               interface{}
	BackgroundPosition             interface{}
	BackgroundPositionX            interface{}
	BackgroundPositionY            interface{}
	BackgroundRepeat               interface{}
	BackgroundRepeatX              interface{}
	BackgroundRepeatY              interface{}
	BackgroundSize                 interface{}
	BaselineShift                  interface{}
	BlockSize                      interface{}
	Border                         interface{}
	BorderBlockEnd                 interface{}
	BorderBlockEndColor            interface{}
	BorderBlockEndStyle            interface{}
	BorderBlockEndWidth            interface{}
	BorderBlockStart               interface{}
	BorderBlockStartColor          interface{}
	BorderBlockStartStyle          interface{}
	BorderBlockStartWidth          interface{}
	BorderBottom                   interface{}
	BorderBottomColor              interface{}
	BorderBottomLeftRadius         interface{}
	BorderBottomRightRadius        interface{}
	BorderBottomStyle              interface{}
	BorderBottomWidth              interface{}
	BorderCollapse                 interface{}
	BorderColor                    interface{}
	BorderImage                    interface{}
	BorderImageOutset              interface{}
	BorderImageRepeat              interface{}
	BorderImageSlice               interface{}
	BorderImageSource              interface{}
	BorderImageWidth               interface{}
	BorderInlineEnd                interface{}
	BorderInlineEndColor           interface{}
	BorderInlineEndStyle           interface{}
	BorderInlineEndWidth           interface{}
	BorderInlineStart              interface{}
	BorderInlineStartColor         interface{}
	BorderInlineStartStyle         interface{}
	BorderInlineStartWidth         interface{}
	BorderLeft                     interface{}
	BorderLeftColor                interface{}
	BorderLeftStyle                interface{}
	BorderRadius                   interface{}
	BorderLeftWidth                interface{}
	BorderRight                    interface{}
	BorderRightColor               interface{}
	BorderRightStyle               interface{}
	BorderRightWidth               interface{}
	BorderSpacing                  interface{}
	BorderStyle                    interface{}
	BorderTop                      interface{}
	BorderTopColor                 interface{}
	BorderTopLeftRadius            interface{}
	BorderTopRightRadius           interface{}
	BorderTopStyle                 interface{}
	BorderTopWidth                 interface{}
	BorderWidth                    interface{}
	Bottom                         interface{}
	BoxShadow                      interface{}
	BoxSizing                      interface{}
	BreakAfter                     interface{}
	BreakBefore                    interface{}
	BreakInside                    interface{}
	BufferedRendering              interface{}
	CaptionSide                    interface{}
	CaretColor                     interface{}
	Clear                          interface{}
	Clip                           interface{}
	ClipPath                       interface{}
	ClipRule                       interface{}
	Color                          interface{}
	ColorInterpolation             interface{}
	ColorInterpolationFilters      interface{}
	ColorRendering                 interface{}
	ColorScheme                    interface{}
	ColumnCount                    interface{}
	ColumnFill                     interface{}
	ColumnGap                      interface{}
	ColumnRule                     interface{}
	ColumnRuleColor                interface{}
	ColumnRuleStyle                interface{}
	ColumnRuleWidth                interface{}
	ColumnSpan                     interface{}
	ColumnWidth                    interface{}
	Columns                        interface{}
	Contain                        interface{}
	ContainIntrinsicSize           interface{}
	Content                        interface{}
	CounterIncrement               interface{}
	CounterReset                   interface{}
	Cursor                         interface{}
	Cx                             interface{}
	Cy                             interface{}
	D                              interface{}
	Direction                      interface{}
	Display                        interface{}
	DominantBaseline               interface{}
	EmptyCells                     interface{}
	Fill                           interface{}
	FillOpacity                    interface{}
	FillRule                       interface{}
	Filter                         interface{}
	Flex                           interface{}
	FlexBasis                      interface{}
	FlexDirection                  interface{}
	FlexFlow                       interface{}
	FlexGrow                       interface{}
	FlexShrink                     interface{}
	FlexWrap                       interface{}
	Float                          interface{}
	FloodColor                     interface{}
	FloodOpacity                   interface{}
	Font                           interface{}
	FontDisplay                    interface{}
	FontFamily                     interface{}
	FontFeatureSettings            interface{}
	FontKerning                    interface{}
	FontOpticalSizing              interface{}
	FontSize                       interface{}
	FontStretch                    interface{}
	FontStyle                      interface{}
	FontVariant                    interface{}
	FontVariantCaps                interface{}
	FontVariantEastAsian           interface{}
	FontVariantLigatures           interface{}
	FontVariantNumeric             interface{}
	FontVariationSettings          interface{}
	FontWeight                     interface{}
	Gap                            interface{}
	Grid                           interface{}
	GridArea                       interface{}
	GridAutoColumns                interface{}
	GridAutoFlow                   interface{}
	GridAutoRows                   interface{}
	GridColumn                     interface{}
	GridColumnEnd                  interface{}
	GridColumnGap                  interface{}
	GridColumnStart                interface{}
	GridGap                        interface{}
	GridRow                        interface{}
	GridRowEnd                     interface{}
	GridRowGap                     interface{}
	GridRowStart                   interface{}
	GridTemplate                   interface{}
	GridTemplateAreas              interface{}
	GridTemplateColumns            interface{}
	GridTemplateRows               interface{}
	Height                         interface{}
	Hyphens                        interface{}
	ImageOrientation               interface{}
	ImageRendering                 interface{}
	InlineSize                     interface{}
	Isolation                      interface{}
	JustifyContent                 interface{}
	JustifyItems                   interface{}
	JustifySelf                    interface{}
	Left                           interface{}
	LetterSpacing                  interface{}
	LightingColor                  interface{}
	LineBreak                      interface{}
	LineHeight                     interface{}
	ListStyle                      interface{}
	ListStyleImage                 interface{}
	ListStylePosition              interface{}
	ListStyleType                  interface{}
	Margin                         interface{}
	MarginBlockEnd                 interface{}
	MarginBlockStart               interface{}
	MarginBottom                   interface{}
	MarginInlineEnd                interface{}
	MarginInlineStart              interface{}
	MarginLeft                     interface{}
	MarginRight                    interface{}
	MarginTop                      interface{}
	Marker                         interface{}
	MarkerEnd                      interface{}
	MarkerMid                      interface{}
	MarkerStart                    interface{}
	Mask                           interface{}
	MaskType                       interface{}
	MaxBlockSize                   interface{}
	MaxHeight                      interface{}
	MaxInlineSize                  interface{}
	MaxWidth                       interface{}
	MaxZoom                        interface{}
	MinBlockSize                   interface{}
	MinHeight                      interface{}
	MinInlineSize                  interface{}
	MinWidth                       interface{}
	MinZoom                        interface{}
	MixBlendMode                   interface{}
	ObjectFit                      interface{}
	ObjectPosition                 interface{}
	Offset                         interface{}
	OffsetDistance                 interface{}
	OffsetPath                     interface{}
	OffsetRotate                   interface{}
	Opacity                        interface{}
	Order                          interface{}
	Orientation                    interface{}
	Orphans                        interface{}
	Outline                        interface{}
	OutlineColor                   interface{}
	OutlineOffset                  interface{}
	OutlineStyle                   interface{}
	OutlineWidth                   interface{}
	Overflow                       interface{}
	OverflowAnchor                 interface{}
	OverflowWrap                   interface{}
	OverflowX                      interface{}
	OverflowY                      interface{}
	OverscrollBehavior             interface{}
	OverscrollBehaviorBlock        interface{}
	OverscrollBehaviorInline       interface{}
	OverscrollBehaviorX            interface{}
	OverscrollBehaviorY            interface{}
	Padding                        interface{}
	PaddingBlockEnd                interface{}
	PaddingBlockStart              interface{}
	PaddingBottom                  interface{}
	PaddingInlineEnd               interface{}
	PaddingInlineStart             interface{}
	PaddingLeft                    interface{}
	PaddingRight                   interface{}
	PaddingTop                     interface{}
	PageBreakAfter                 interface{}
	PageBreakBefore                interface{}
	PageBreakInside                interface{}
	PaintOrder                     interface{}
	Perspective                    interface{}
	PerspectiveOrigin              interface{}
	PlaceContent                   interface{}
	PlaceItems                     interface{}
	PlaceSelf                      interface{}
	PointerEvents                  interface{}
	Position                       interface{}
	Quotes                         interface{}
	R                              interface{}
	Resize                         interface{}
	Right                          interface{}
	RowGap                         interface{}
	Rx                             interface{}
	Ry                             interface{}
	ScrollBehavior                 interface{}
	ScrollMargin                   interface{}
	ScrollMarginBlock              interface{}
	ScrollMarginBlockEnd           interface{}
	ScrollMarginBlockStart         interface{}
	ScrollMarginBottom             interface{}
	ScrollMarginInline             interface{}
	ScrollMarginInlineEnd          interface{}
	ScrollMarginInlineStart        interface{}
	ScrollMarginLeft               interface{}
	ScrollMarginRight              interface{}
	ScrollMarginTop                interface{}
	ScrollPadding                  interface{}
	ScrollPaddingBlock             interface{}
	ScrollPaddingBlockEnd          interface{}
	ScrollPaddingBlockStart        interface{}
	ScrollPaddingBottom            interface{}
	ScrollPaddingInline            interface{}
	ScrollPaddingInlineEnd         interface{}
	ScrollPaddingInlineStart       interface{}
	ScrollPaddingLeft              interface{}
	ScrollPaddingRight             interface{}
	ScrollPaddingTop               interface{}
	ScrollSnapAlign                interface{}
	ScrollSnapStop                 interface{}
	ScrollSnapType                 interface{}
	ShapeImageThreshold            interface{}
	ShapeMargin                    interface{}
	ShapeOutside                   interface{}
	ShapeRendering                 interface{}
	Size                           interface{}
	Speak                          interface{}
	Src                            interface{}
	StopColor                      interface{}
	StopOpacity                    interface{}
	Stroke                         interface{}
	StrokeDasharray                interface{}
	StrokeDashoffset               interface{}
	StrokeLinecap                  interface{}
	StrokeLinejoin                 interface{}
	StrokeMiterlimit               interface{}
	StrokeOpacity                  interface{}
	StrokeWidth                    interface{}
	TabSize                        interface{}
	TableLayout                    interface{}
	TextAlign                      interface{}
	TextAlignLast                  interface{}
	TextAnchor                     interface{}
	TextCombineUpright             interface{}
	TextDecoration                 interface{}
	TextDecorationColor            interface{}
	TextDecorationLine             interface{}
	TextDecorationSkipInk          interface{}
	TextDecorationStyle            interface{}
	TextIndent                     interface{}
	TextOrientation                interface{}
	TextOverflow                   interface{}
	TextRendering                  interface{}
	TextShadow                     interface{}
	TextSizeAdjust                 interface{}
	TextTransform                  interface{}
	TextUnderlinePosition          interface{}
	Top                            interface{}
	TouchAction                    interface{}
	Transform                      interface{}
	TransformBox                   interface{}
	TransformOrigin                interface{}
	TransformStyle                 interface{}
	Transition                     interface{}
	TransitionDelay                interface{}
	TransitionDuration             interface{}
	TransitionProperty             interface{}
	TransitionTimingFunction       interface{}
	UnicodeBidi                    interface{}
	UnicodeRange                   interface{}
	UserSelect                     interface{}
	UserZoom                       interface{}
	VectorEffect                   interface{}
	VerticalAlign                  interface{}
	Visibility                     interface{}
	WebkitAlignContent             interface{}
	WebkitAlignItems               interface{}
	WebkitAlignSelf                interface{}
	WebkitAnimation                interface{}
	WebkitAnimationDelay           interface{}
	WebkitAnimationDirection       interface{}
	WebkitAnimationDuration        interface{}
	WebkitAnimationFillMode        interface{}
	WebkitAnimationIterationCount  interface{}
	WebkitAnimationName            interface{}
	WebkitAnimationPlayState       interface{}
	WebkitAnimationTimingFunction  interface{}
	WebkitAppRegion                interface{}
	WebkitAppearance               interface{}
	WebkitBackfaceVisibility       interface{}
	WebkitBackgroundClip           interface{}
	WebkitBackgroundOrigin         interface{}
	WebkitBackgroundSize           interface{}
	WebkitBorderAfter              interface{}
	WebkitBorderAfterColor         interface{}
	WebkitBorderAfterStyle         interface{}
	WebkitBorderAfterWidth         interface{}
	WebkitBorderBefore             interface{}
	WebkitBorderBeforeColor        interface{}
	WebkitBorderBeforeStyle        interface{}
	WebkitBorderBeforeWidth        interface{}
	WebkitBorderBottomLeftRadius   interface{}
	WebkitBorderBottomRightRadius  interface{}
	WebkitBorderEnd                interface{}
	WebkitBorderEndColor           interface{}
	WebkitBorderEndStyle           interface{}
	WebkitBorderEndWidth           interface{}
	WebkitBorderHorizontalSpacing  interface{}
	WebkitBorderImage              interface{}
	WebkitBorderRadius             interface{}
	WebkitBorderStart              interface{}
	WebkitBorderStartColor         interface{}
	WebkitBorderStartStyle         interface{}
	WebkitBorderStartWidth         interface{}
	WebkitBorderTopLeftRadius      interface{}
	WebkitBorderTopRightRadius     interface{}
	WebkitBorderVerticalSpacing    interface{}
	WebkitBoxAlign                 interface{}
	WebkitBoxDecorationBreak       interface{}
	WebkitBoxDirection             interface{}
	WebkitBoxFlex                  interface{}
	WebkitBoxOrdinalGroup          interface{}
	WebkitBoxOrient                interface{}
	WebkitBoxPack                  interface{}
	WebkitBoxReflect               interface{}
	WebkitBoxShadow                interface{}
	WebkitBoxSizing                interface{}
	WebkitClipPath                 interface{}
	WebkitColumnBreakAfter         interface{}
	WebkitColumnBreakBefore        interface{}
	WebkitColumnBreakInside        interface{}
	WebkitColumnCount              interface{}
	WebkitColumnGap                interface{}
	WebkitColumnRule               interface{}
	WebkitColumnRuleColor          interface{}
	WebkitColumnRuleStyle          interface{}
	WebkitColumnRuleWidth          interface{}
	WebkitColumnSpan               interface{}
	WebkitColumnWidth              interface{}
	WebkitColumns                  interface{}
	WebkitFilter                   interface{}
	WebkitFlex                     interface{}
	WebkitFlexBasis                interface{}
	WebkitFlexDirection            interface{}
	WebkitFlexFlow                 interface{}
	WebkitFlexGrow                 interface{}
	WebkitFlexShrink               interface{}
	WebkitFlexWrap                 interface{}
	WebkitFontFeatureSettings      interface{}
	WebkitFontSizeDelta            interface{}
	WebkitFontSmoothing            interface{}
	WebkitHighlight                interface{}
	WebkitHyphenateCharacter       interface{}
	WebkitJustifyContent           interface{}
	WebkitLineBreak                interface{}
	WebkitLineClamp                interface{}
	WebkitLocale                   interface{}
	WebkitLogicalHeight            interface{}
	WebkitLogicalWidth             interface{}
	WebkitMarginAfter              interface{}
	WebkitMarginBefore             interface{}
	WebkitMarginEnd                interface{}
	WebkitMarginStart              interface{}
	WebkitMask                     interface{}
	WebkitMaskBoxImage             interface{}
	WebkitMaskBoxImageOutset       interface{}
	WebkitMaskBoxImageRepeat       interface{}
	WebkitMaskBoxImageSlice        interface{}
	WebkitMaskBoxImageSource       interface{}
	WebkitMaskBoxImageWidth        interface{}
	WebkitMaskClip                 interface{}
	WebkitMaskComposite            interface{}
	WebkitMaskImage                interface{}
	WebkitMaskOrigin               interface{}
	WebkitMaskPosition             interface{}
	WebkitMaskPositionX            interface{}
	WebkitMaskPositionY            interface{}
	WebkitMaskRepeat               interface{}
	WebkitMaskRepeatX              interface{}
	WebkitMaskRepeatY              interface{}
	WebkitMaskSize                 interface{}
	WebkitMaxLogicalHeight         interface{}
	WebkitMaxLogicalWidth          interface{}
	WebkitMinLogicalHeight         interface{}
	WebkitMinLogicalWidth          interface{}
	WebkitOpacity                  interface{}
	WebkitOrder                    interface{}
	WebkitPaddingAfter             interface{}
	WebkitPaddingBefore            interface{}
	WebkitPaddingEnd               interface{}
	WebkitPaddingStart             interface{}
	WebkitPerspective              interface{}
	WebkitPerspectiveOrigin        interface{}
	WebkitPerspectiveOriginX       interface{}
	WebkitPerspectiveOriginY       interface{}
	WebkitPrintColorAdjust         interface{}
	WebkitRtlOrdering              interface{}
	WebkitRubyPosition             interface{}
	WebkitShapeImageThreshold      interface{}
	WebkitShapeMargin              interface{}
	WebkitShapeOutside             interface{}
	WebkitTapHighlightColor        interface{}
	WebkitTextCombine              interface{}
	WebkitTextDecorationsInEffect  interface{}
	WebkitTextEmphasis             interface{}
	WebkitTextEmphasisColor        interface{}
	WebkitTextEmphasisPosition     interface{}
	WebkitTextEmphasisStyle        interface{}
	WebkitTextFillColor            interface{}
	WebkitTextOrientation          interface{}
	WebkitTextSecurity             interface{}
	WebkitTextSizeAdjust           interface{}
	WebkitTextStroke               interface{}
	WebkitTextStrokeColor          interface{}
	WebkitTextStrokeWidth          interface{}
	WebkitTransform                interface{}
	WebkitTransformOrigin          interface{}
	WebkitTransformOriginX         interface{}
	WebkitTransformOriginY         interface{}
	WebkitTransformOriginZ         interface{}
	WebkitTransformStyle           interface{}
	WebkitTransition               interface{}
	WebkitTransitionDelay          interface{}
	WebkitTransitionDuration       interface{}
	WebkitTransitionProperty       interface{}
	WebkitTransitionTimingFunction interface{}
	WebkitUserDrag                 interface{}
	WebkitUserModify               interface{}
	WebkitUserSelect               interface{}
	WebkitWritingMode              interface{}
	WhiteSpace                     interface{}
	Widows                         interface{}
	Width                          interface{}
	WillChange                     interface{}
	WordBreak                      interface{}
	WordSpacing                    interface{}
	WordWrap                       interface{}
	WritingMode                    interface{}
	X                              interface{}
	Y                              interface{}
	ZIndex                         interface{}
	Zoom                           interface{}
}

func (s *Style) String() string {
	values := toMap(s)
	var attrs []string

	for key, value := range values {
		attr := s.property(key)
		strValue := ""

		switch value.(type) {
		case []string:
			strValue = strings.Join(value.([]string), " ")
			break

		case map[string]string:
			strValue = ""
			for mk, mv := range value.(map[string]string) {
				strValue += mk + "(" + mv + ") "
			}
			break

		default:
			strValue = cast.ToString(value)
			break
		}

		if strValue != "" {
			attrs = append(attrs, attr+":"+strValue)
		}
	}

	return strings.Join(attrs, ";")
}

func (s *Style) property(key string) string {
	return strcase.ToKebab(key)
}
