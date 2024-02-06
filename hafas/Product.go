package hafas

type Product struct {
	ID      string
	Name    string
	Bitmask uint8
}

type ProductID uint

// Enum-like constants for Product
const (
	ProductSuburban ProductID = iota + 1
	ProductSubway
	ProductTram
	ProductBus
	ProductFerry
	ProductExpress
	ProductRegional
)

// MaxProductsFilterBitmask represents the combined bitmask of all products.
// It is used to validate that a given bitmask does not exceed the combined value of all available products.
var MaxProductsFilterBitmask uint8

// Products maps Product constants to their names and bitmasks
var Products = map[ProductID]Product{
	ProductSuburban: {"S", "S-Bahn", 1},
	ProductSubway:   {"U", "U-Bahn", 2},
	ProductTram:     {"T", "Tram", 4},
	ProductBus:      {"B", "Bus", 8},
	ProductFerry:    {"F", "Fähre", 16},
	ProductExpress:  {"E", "IC/ICE", 32},
	ProductRegional: {"R", "RB/RE", 64},
}

func init() {
	for _, p := range Products {
		MaxProductsFilterBitmask |= p.Bitmask
	}
}

// ProductsFilter is a bitmask representing a set of products
type ProductsFilter struct {
	bitmask uint8
}

func NewProductFilter() *ProductsFilter {
	return &ProductsFilter{
		bitmask: 127,
	}
}

func (f *ProductsFilter) AddProduct(id ProductID) *ProductsFilter {
	if f.bitmask >= MaxProductsFilterBitmask {
		return f
	}
	if product, ok := Products[id]; ok {
		f.bitmask |= product.Bitmask
	}
	return f
}

func (f *ProductsFilter) Build() uint8 {
	return f.bitmask
}
