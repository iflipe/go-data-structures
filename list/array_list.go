package list

import "errors"

type ArrayList struct {
	values   []int
	inserted int
}

func (list *ArrayList) Init(size int) error {
	if size > 0 {
		list.values = make([]int, size)
		return nil
	} else {
		return errors.New("can't init arraylist with size <= 0")
	}
}

/**
 * Duplica o vetor.
 */
func (list *ArrayList) doubleArray() {
	curSize := len(list.values)
	doubledValues := make([]int, 2*curSize)

	for i := 0; i < curSize; i++ {
		doubledValues[i] = list.values[i]
	}
	list.values = doubledValues
}

/**
 * Adiciona sempre da esquerda para a direita,
 * após o último elemento inserido anteriormente.
 *
 * Melhor caso: Ômega(1)
 *   Just: não precisa duplicar vetor, nem varrer o vetor
 *         do início para encontrar o endereço a ser Add
 * Pior caso: O(n)
 *   Just: duplicar o vetor requer copiar os elementos
 *         para o novo vetor, causando um custo computa-
 *         cional proporcional ao tamanho do vetor (n)
 */
func (list *ArrayList) Add(val int) {
	//verificar se array está lotado
	if list.inserted >= len(list.values) {
		list.doubleArray()
	}
	list.values[list.inserted] = val
	list.inserted++
}

/**
 * Adiciona elemento na posição especificada, deslocando
 * os elementos à direita dessa posição.
 *
 * Melhor caso: Ômega(1)
 *   Just: não precisa duplicar vetor, nem varrer o vetor
 *         do início para encontrar o endereço a ser Add
 * Pior caso: O(n)
 *   Just: 1) duplicar o vetor requer copiar os elementos
 *         para o novo vetor, causando um custo computa-
 *         cional proporcional ao tamanho do vetor (n)
 *         2) adicionar no início requer deslocar os n
 *            elementos do vetor para a direita
 */
func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.inserted {
		if list.inserted == len(list.values) {
			list.doubleArray()
		}
		for i := list.inserted; i > index; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[index] = value
		list.inserted++
		return nil
	} else {
		if index < 0 {
			return errors.New("can't add in arraylist on negative index")
		} else {
			return errors.New("can't add in arraylist on index out of upper bounds")
		}
	}
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < list.inserted {
		for i := index; i < list.inserted-1; i++ {
			list.values[i] = list.values[i+1]
		}
		list.inserted--
		return nil
	} else {
		if index < 0 {
			return errors.New("can't remove from arraylist on negative index")
		} else {
			return errors.New("can't remove from arraylist on index out of upper bounds")
		}
	}
}

func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.values[index], nil
	} else {
		if index < 0 {
			return index, errors.New("can't get value from arraylist on negative index")
		} else {
			return index, errors.New("can't get value from arraylist on index out of upper bounds")
		}
	}
}

func (list *ArrayList) Set(value, index int) error {
	if index >= 0 && index < list.inserted {
		list.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("can't set value on arraylist on index < 0")
		} else {
			return errors.New("can't set value on arraylist on index >= arraylist.size")
		}
	}
}

func (list *ArrayList) Size() int {
	return list.inserted
}

func (list *ArrayList) Reverse() {
	for i := 0; i < len(list.values)/2; i++ {
		list.values[i], list.values[len(list.values)-i-1] = list.values[len(list.values)-i-1], list.values[i]
	}
}

func (list *ArrayList) BSeach(value int, init int, end int) int {
	if init > end {
		return -1
	}
	mid := (init + end) / 2
	if list.values[mid] == value {
		return mid
	}
	if list.values[mid] < value {
		return list.BSeach(value, mid+1, end)
	} else {
		return list.BSeach(value, init, mid-1)
	}
}

func (list *ArrayList) LSearch(value int) int {
	for i := 0; i < list.inserted; i++ {
		if list.values[i] == value {
			return i
		}
	}
	return -1
}
func (list *ArrayList) BubbleSort() {
	has_swapped := false
	for i := 0; i < list.Size()-1; i++ {
		has_swapped = false
		for j := 0; j < list.Size()-i-1; j++ {
			if list.values[j] > list.values[j+1] {
				list.values[j], list.values[j+1] = list.values[j+1], list.values[j]
				has_swapped = true
			}
		}
		if !has_swapped {
			return
		}
	}
}

/*
func (list *ArrayList) InsertionSort() {
	for i := 1; i < list.Size(); i++ {
		for j := i; j > 0 && list.values[j-1] > list.values[j]; j-- {
			list.values[j-1], list.values[j] = list.values[j], list.values[j-1]
		}
	}
}
*/

func (list *ArrayList) InsertionSort() {
	for i := 1; i < list.Size(); i++ {
		j := i
		tmp := list.values[i]
		for j > 0 && list.values[j-1] > tmp {
			list.values[j] = list.values[j-1]
			j--
		}
		list.values[j] = tmp
	}
}
