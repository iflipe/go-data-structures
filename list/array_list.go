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
