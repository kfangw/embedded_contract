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

package operation

import (
	"io"

	"errors"
	"tc/common/address"
	"tc/common/serialization"
)

type OperationInvoke struct {
	Version byte
	Address address.Address
	Method  string
	Args    []byte
}

// Serialize contract
func (oi *OperationInvoke) Serialize(w io.Writer) error {
	if err := serialization.WriteByte(w, oi.Version); err != nil {
		return errors.New("[OperationInvoke] Version serialize error!")
	}
	if err := oi.Address.Serialize(w); err != nil {
		return errors.New("[OperationInvoke] Address serialize error!")
	}
	if err := serialization.WriteVarBytes(w, []byte(oi.Method)); err != nil {
		return errors.New("[OperationInvoke] Method serialize error!")
	}
	if err := serialization.WriteVarBytes(w, oi.Args); err != nil {
		return errors.New("[OperationInvoke] Args serialize error!")
	}
	return nil
}

// Deserialize contract
func (oi *OperationInvoke) Deserialize(r io.Reader) error {
	var err error
	oi.Version, err = serialization.ReadByte(r)
	if err != nil {
		return errors.New("[OperationInvoke] Version deserialize error!")
	}

	if err := oi.Address.Deserialize(r); err != nil {
		return errors.New("[OperationInvoke] Address deserialize error!")
	}

	method, err := serialization.ReadVarBytes(r)
	if err != nil {
		return errors.New("[OperationInvoke] Method deserialize error!")
	}
	oi.Method = string(method)

	oi.Args, err = serialization.ReadVarBytes(r)
	if err != nil {
		return errors.New("[OperationInvoke] Args deserialize error!")
	}
	return nil
}
