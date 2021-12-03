// Code generated by entc, DO NOT EDIT.

package book

import (
	"back/repository/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// SurplusCatch applies equality check predicate on the "surplus_catch" field. It's identical to SurplusCatchEQ.
func SurplusCatch(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSurplusCatch), v))
	})
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// Describe applies equality check predicate on the "describe" field. It's identical to DescribeEQ.
func Describe(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescribe), v))
	})
}

// Ebook applies equality check predicate on the "ebook" field. It's identical to EbookEQ.
func Ebook(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEbook), v))
	})
}

// Cover applies equality check predicate on the "cover" field. It's identical to CoverEQ.
func Cover(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCover), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// SurplusCatchEQ applies the EQ predicate on the "surplus_catch" field.
func SurplusCatchEQ(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSurplusCatch), v))
	})
}

// SurplusCatchNEQ applies the NEQ predicate on the "surplus_catch" field.
func SurplusCatchNEQ(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSurplusCatch), v))
	})
}

// SurplusCatchIn applies the In predicate on the "surplus_catch" field.
func SurplusCatchIn(vs ...int) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSurplusCatch), v...))
	})
}

// SurplusCatchNotIn applies the NotIn predicate on the "surplus_catch" field.
func SurplusCatchNotIn(vs ...int) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSurplusCatch), v...))
	})
}

// SurplusCatchGT applies the GT predicate on the "surplus_catch" field.
func SurplusCatchGT(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSurplusCatch), v))
	})
}

// SurplusCatchGTE applies the GTE predicate on the "surplus_catch" field.
func SurplusCatchGTE(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSurplusCatch), v))
	})
}

// SurplusCatchLT applies the LT predicate on the "surplus_catch" field.
func SurplusCatchLT(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSurplusCatch), v))
	})
}

// SurplusCatchLTE applies the LTE predicate on the "surplus_catch" field.
func SurplusCatchLTE(v int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSurplusCatch), v))
	})
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthor), v))
	})
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthor), v...))
	})
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthor), v...))
	})
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthor), v))
	})
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthor), v))
	})
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthor), v))
	})
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthor), v))
	})
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthor), v))
	})
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthor), v))
	})
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthor), v))
	})
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthor), v))
	})
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthor), v))
	})
}

// DescribeEQ applies the EQ predicate on the "describe" field.
func DescribeEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescribe), v))
	})
}

// DescribeNEQ applies the NEQ predicate on the "describe" field.
func DescribeNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescribe), v))
	})
}

// DescribeIn applies the In predicate on the "describe" field.
func DescribeIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescribe), v...))
	})
}

// DescribeNotIn applies the NotIn predicate on the "describe" field.
func DescribeNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescribe), v...))
	})
}

// DescribeGT applies the GT predicate on the "describe" field.
func DescribeGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescribe), v))
	})
}

// DescribeGTE applies the GTE predicate on the "describe" field.
func DescribeGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescribe), v))
	})
}

// DescribeLT applies the LT predicate on the "describe" field.
func DescribeLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescribe), v))
	})
}

// DescribeLTE applies the LTE predicate on the "describe" field.
func DescribeLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescribe), v))
	})
}

// DescribeContains applies the Contains predicate on the "describe" field.
func DescribeContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescribe), v))
	})
}

// DescribeHasPrefix applies the HasPrefix predicate on the "describe" field.
func DescribeHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescribe), v))
	})
}

// DescribeHasSuffix applies the HasSuffix predicate on the "describe" field.
func DescribeHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescribe), v))
	})
}

// DescribeEqualFold applies the EqualFold predicate on the "describe" field.
func DescribeEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescribe), v))
	})
}

// DescribeContainsFold applies the ContainsFold predicate on the "describe" field.
func DescribeContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescribe), v))
	})
}

// EbookEQ applies the EQ predicate on the "ebook" field.
func EbookEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEbook), v))
	})
}

// EbookNEQ applies the NEQ predicate on the "ebook" field.
func EbookNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEbook), v))
	})
}

// EbookIn applies the In predicate on the "ebook" field.
func EbookIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEbook), v...))
	})
}

// EbookNotIn applies the NotIn predicate on the "ebook" field.
func EbookNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEbook), v...))
	})
}

// EbookGT applies the GT predicate on the "ebook" field.
func EbookGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEbook), v))
	})
}

// EbookGTE applies the GTE predicate on the "ebook" field.
func EbookGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEbook), v))
	})
}

// EbookLT applies the LT predicate on the "ebook" field.
func EbookLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEbook), v))
	})
}

// EbookLTE applies the LTE predicate on the "ebook" field.
func EbookLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEbook), v))
	})
}

// EbookContains applies the Contains predicate on the "ebook" field.
func EbookContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEbook), v))
	})
}

// EbookHasPrefix applies the HasPrefix predicate on the "ebook" field.
func EbookHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEbook), v))
	})
}

// EbookHasSuffix applies the HasSuffix predicate on the "ebook" field.
func EbookHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEbook), v))
	})
}

// EbookEqualFold applies the EqualFold predicate on the "ebook" field.
func EbookEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEbook), v))
	})
}

// EbookContainsFold applies the ContainsFold predicate on the "ebook" field.
func EbookContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEbook), v))
	})
}

// CoverEQ applies the EQ predicate on the "cover" field.
func CoverEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCover), v))
	})
}

// CoverNEQ applies the NEQ predicate on the "cover" field.
func CoverNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCover), v))
	})
}

// CoverIn applies the In predicate on the "cover" field.
func CoverIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCover), v...))
	})
}

// CoverNotIn applies the NotIn predicate on the "cover" field.
func CoverNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCover), v...))
	})
}

// CoverGT applies the GT predicate on the "cover" field.
func CoverGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCover), v))
	})
}

// CoverGTE applies the GTE predicate on the "cover" field.
func CoverGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCover), v))
	})
}

// CoverLT applies the LT predicate on the "cover" field.
func CoverLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCover), v))
	})
}

// CoverLTE applies the LTE predicate on the "cover" field.
func CoverLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCover), v))
	})
}

// CoverContains applies the Contains predicate on the "cover" field.
func CoverContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCover), v))
	})
}

// CoverHasPrefix applies the HasPrefix predicate on the "cover" field.
func CoverHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCover), v))
	})
}

// CoverHasSuffix applies the HasSuffix predicate on the "cover" field.
func CoverHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCover), v))
	})
}

// CoverEqualFold applies the EqualFold predicate on the "cover" field.
func CoverEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCover), v))
	})
}

// CoverContainsFold applies the ContainsFold predicate on the "cover" field.
func CoverContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCover), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		p(s.Not())
	})
}