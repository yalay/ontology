/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package neovm

func opCat(e *ExecutionEngine) (VMState, error) {
	b2 := PopByteArray(e)
	b1 := PopByteArray(e)
	r := Concat(b1, b2)
	PushData(e, r)
	return NONE, nil
}

func opSubStr(e *ExecutionEngine) (VMState, error) {
	count := PopInt(e)
	index := PopInt(e)
	arr := PopByteArray(e)
	b := arr[index : index+count]
	PushData(e, b)
	return NONE, nil
}

func opLeft(e *ExecutionEngine) (VMState, error) {
	count := PopInt(e)
	s := PopByteArray(e)
	b := s[:count]
	PushData(e, b)
	return NONE, nil
}

func opRight(e *ExecutionEngine) (VMState, error) {
	count := PopInt(e)
	arr := PopByteArray(e)
	b := arr[len(arr)-count:]
	PushData(e, b)
	return NONE, nil
}

func opSize(e *ExecutionEngine) (VMState, error) {
	x := Peek(e).GetStackItem()
	PushData(e, len(x.GetByteArray()))
	return NONE, nil
}
