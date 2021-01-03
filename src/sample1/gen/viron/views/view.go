// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Viron views
//
// Command:
// $ goa gen sample1/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// VironAuthtypeCollection is the viewed result type that is projected based on
// a view.
type VironAuthtypeCollection struct {
	// Type to project
	Projected VironAuthtypeCollectionView
	// View to render
	View string
}

// VironMenu is the viewed result type that is projected based on a view.
type VironMenu struct {
	// Type to project
	Projected *VironMenuView
	// View to render
	View string
}

// VironAuthtypeCollectionView is a type that runs validations on a projected
// type.
type VironAuthtypeCollectionView []*VironAuthtypeView

// VironAuthtypeView is a type that runs validations on a projected type.
type VironAuthtypeView struct {
	// type name
	Type *string
	// provider name
	Provider *string
	// url
	URL *string
	// request method to submit this auth
	Method *string
}

// VironMenuView is a type that runs validations on a projected type.
type VironMenuView struct {
	Name      *string
	Thumbnail *string
	Tags      []string
	Color     *string
	Theme     *string
	Pages     []*VironPageView
	Sections  []*VironSectionView
}

// VironPageView is a type that runs validations on a projected type.
type VironPageView struct {
	ID         *string
	Name       *string
	Section    *string
	Group      *string
	Components []*VironComponentView
}

// VironComponentView is a type that runs validations on a projected type.
type VironComponentView struct {
	Name           *string
	Style          *string
	Unit           *string
	Actions        []string
	API            *VironAPIView
	Pagination     *bool
	Primary        *string
	TableLabels    []string
	Query          []*VironQueryView
	AutoRefreshSec *int32
}

// VironAPIView is a type that runs validations on a projected type.
type VironAPIView struct {
	Method *string
	Path   *string
}

// VironQueryView is a type that runs validations on a projected type.
type VironQueryView struct {
	Key  *string
	Type *string
}

// VironSectionView is a type that runs validations on a projected type.
type VironSectionView struct {
	ID    *string
	Label *string
}

var (
	// VironAuthtypeCollectionMap is a map of attribute names in result type
	// VironAuthtypeCollection indexed by view name.
	VironAuthtypeCollectionMap = map[string][]string{
		"default": []string{
			"type",
			"provider",
			"url",
			"method",
		},
	}
	// VironMenuMap is a map of attribute names in result type VironMenu indexed by
	// view name.
	VironMenuMap = map[string][]string{
		"default": []string{
			"name",
			"thumbnail",
			"tags",
			"color",
			"theme",
			"pages",
		},
	}
	// VironAuthtypeMap is a map of attribute names in result type VironAuthtype
	// indexed by view name.
	VironAuthtypeMap = map[string][]string{
		"default": []string{
			"type",
			"provider",
			"url",
			"method",
		},
	}
	// VironPageMap is a map of attribute names in result type VironPage indexed by
	// view name.
	VironPageMap = map[string][]string{
		"default": []string{
			"section",
			"group",
			"id",
			"name",
			"components",
		},
	}
	// VironComponentMap is a map of attribute names in result type VironComponent
	// indexed by view name.
	VironComponentMap = map[string][]string{
		"default": []string{
			"name",
			"style",
			"unit",
			"actions",
			"api",
			"pagination",
			"primary",
			"table_labels",
			"query",
			"auto_refresh_sec",
		},
	}
	// VironAPIMap is a map of attribute names in result type VironAPI indexed by
	// view name.
	VironAPIMap = map[string][]string{
		"default": []string{
			"method",
			"path",
		},
	}
	// VironQueryMap is a map of attribute names in result type VironQuery indexed
	// by view name.
	VironQueryMap = map[string][]string{
		"default": []string{
			"key",
			"type",
		},
	}
	// VironSectionMap is a map of attribute names in result type VironSection
	// indexed by view name.
	VironSectionMap = map[string][]string{
		"default": []string{
			"id",
			"label",
		},
	}
)

// ValidateVironAuthtypeCollection runs the validations defined on the viewed
// result type VironAuthtypeCollection.
func ValidateVironAuthtypeCollection(result VironAuthtypeCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateVironAuthtypeCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateVironMenu runs the validations defined on the viewed result type
// VironMenu.
func ValidateVironMenu(result *VironMenu) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateVironMenuView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateVironAuthtypeCollectionView runs the validations defined on
// VironAuthtypeCollectionView using the "default" view.
func ValidateVironAuthtypeCollectionView(result VironAuthtypeCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateVironAuthtypeView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVironAuthtypeView runs the validations defined on VironAuthtypeView
// using the "default" view.
func ValidateVironAuthtypeView(result *VironAuthtypeView) (err error) {
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Provider == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provider", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	if result.Method == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("method", "result"))
	}
	return
}

// ValidateVironMenuView runs the validations defined on VironMenuView using
// the "default" view.
func ValidateVironMenuView(result *VironMenuView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Pages == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pages", "result"))
	}
	if result.Color != nil {
		if !(*result.Color == "purple" || *result.Color == "blue" || *result.Color == "green" || *result.Color == "yellow" || *result.Color == "red" || *result.Color == "gray" || *result.Color == "black" || *result.Color == "white") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.color", *result.Color, []interface{}{"purple", "blue", "green", "yellow", "red", "gray", "black", "white"}))
		}
	}
	if result.Theme != nil {
		if !(*result.Theme == "standard" || *result.Theme == "midnight" || *result.Theme == "terminal") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.theme", *result.Theme, []interface{}{"standard", "midnight", "terminal"}))
		}
	}
	for _, e := range result.Pages {
		if e != nil {
			if err2 := ValidateVironPageView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateVironPageView runs the validations defined on VironPageView using
// the "default" view.
func ValidateVironPageView(result *VironPageView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Section == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("section", "result"))
	}
	if result.Components == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("components", "result"))
	}
	for _, e := range result.Components {
		if e != nil {
			if err2 := ValidateVironComponentView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateVironComponentView runs the validations defined on
// VironComponentView using the "default" view.
func ValidateVironComponentView(result *VironComponentView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Style == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("style", "result"))
	}
	if result.Style != nil {
		if !(*result.Style == "number" || *result.Style == "table" || *result.Style == "graph-bar" || *result.Style == "graph-scatterplot" || *result.Style == "graph-line" || *result.Style == "graph-horizontal-bar" || *result.Style == "graph-stacked-bar" || *result.Style == "graph-horizontal-stacked-bar" || *result.Style == "graph-stacked-area") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.style", *result.Style, []interface{}{"number", "table", "graph-bar", "graph-scatterplot", "graph-line", "graph-horizontal-bar", "graph-stacked-bar", "graph-horizontal-stacked-bar", "graph-stacked-area"}))
		}
	}
	for _, e := range result.Query {
		if e != nil {
			if err2 := ValidateVironQueryView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.API != nil {
		if err2 := ValidateVironAPIView(result.API); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVironAPIView runs the validations defined on VironAPIView using the
// "default" view.
func ValidateVironAPIView(result *VironAPIView) (err error) {
	if result.Method == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("method", "result"))
	}
	if result.Path == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("path", "result"))
	}
	return
}

// ValidateVironQueryView runs the validations defined on VironQueryView using
// the "default" view.
func ValidateVironQueryView(result *VironQueryView) (err error) {
	if result.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	return
}

// ValidateVironSectionView runs the validations defined on VironSectionView
// using the "default" view.
func ValidateVironSectionView(result *VironSectionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "result"))
	}
	return
}