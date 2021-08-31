package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBuiltinRwMapDeleteParalell(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(100000000)
			buildinRwMapDelete(k)
		}
	})
}

func BenchmarkBuiltinMapDeleteParalell(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(100000000)
			buildinMapDelete(k)
		}
	})
}

func BenchmarkBuiltinSyncMapDeleteParalell(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(100000000)
			buildinSyncMapDelete(k)
		}
	})
}

func BenchmarkBuiltinRwMapStoreParalell(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(100000000)
			buildinRwMapStore(k, k)
		}
	})
}

func BenchmarkBuiltinSyncMapStoreParalell(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(10000000)
			buildinSyncMapStore(k, k)
		}
	})
}

func BenchmarkBuiltinMapStoreParalell(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(100000000)
			buildinMapStore(k, k)
		}
	})
}

func BenchmarkBuiltinMapLookUp(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(1000000000)
			buildinMapLookUp(k)
		}
	})
}
func BenchmarkBuiltinSyncMapLookUp(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(1000000000)
			buildinSyncMapLookUp(k)
		}
	})
}
func BenchmarkBuiltinRWMapLookUp(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(1000000000)
			buildinRwMapLookup(k)
		}
	})
}