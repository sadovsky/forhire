package item

type ItemClass int
const (
  Armor ItemClass = iota
)

type Item struct {
  Name string
  Type ItemClass
}
