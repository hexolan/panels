// Copyright 2023 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Document, Schema, model } from "mongoose";
import uniqueValidator from "mongoose-unique-validator";

const userSchema = new Schema(
  {
    username: { type: String, required: true, lowercase: true, unique: true },
    isAdmin: { type: Boolean, required: false, default: false }
  },
  { 
    timestamps: true
  }
);

userSchema.plugin(uniqueValidator);

interface IUser extends Document {
  username: string;
  isAdmin?: boolean;
  createdAt: Date;
  updatedAt: Date;
}

const User = model<IUser>("User", userSchema);

export { User, IUser }