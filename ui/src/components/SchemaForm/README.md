---
sidebar_position: 0
---
# Schema Form

## Introduction

A React component capable of building HTML forms out of a [JSON schema](https://json-schema.org/understanding-json-schema/index.html).

## Usage

```tsx
import React, { useState } from 'react';

import { SchemaForm, initFormData, JSONSchema, UISchema } from '@/components';

const schema: JSONSchema = {
  title: 'General',
  properties: {
    name: {
      type: 'string',
      title: 'Name',
    },
    age: {
      type: 'number',
      title: 'Age',
    },
    sex: {
      type: 'string',
      title: 'sex',
      enum: [1, 2],
      enumNames: ['male', 'female'],
    },
  },
};

const uiSchema: UISchema = {
  name: {
    'ui:widget': 'input',
  },
  age: {
    'ui:widget': 'input',
    'ui:options': {
      type: 'number',
    },
  },
  sex: {
    'ui:widget': 'radio',
  },
};

const Form = () => {
  const [formData, setFormData] = useState(initFormData(schema));

  const handleChange = (data) => {
    setFormData(data);
  };

  return (
    <SchemaForm
      schema={schema}
      uiSchema={uiSchema}
      formData={formData}
      onChange={handleChange}
    />
  );
};

export default Form;

```

## Props

| Property | Description                              | Type                                  | Default |
| -------- | ---------------------------------------- | ------------------------------------- | ------- |
| schema   | Describe the form structure with schema  | [JSONSchema](#json-schema)            | -       |
| uiSchema | Describe the properties of the field     | [UISchema](#uischema)                 | -       |
| formData | Describe form data                       | [FormData](#formdata)                 | -       |
| onChange | Callback function when form data changes | (data: [FormData](#formdata)) => void | -       |
| onSubmit | Callback function when form is submitted | (data: React.FormEvent) => void       | -       |

## Types Definition
### JSONSchema

```ts
export interface JSONSchema {
  title: string;
  description?: string;
  required?: string[];
  properties: {
    [key: string]: {
      type: 'string' | 'boolean' | 'number';
      title: string;
      label?: string;
      description?: string;
      enum?: Array<string | boolean | number>;
      enumNames?: string[];
      default?: string | boolean | number;
    };
  };
}
```

### UIOptions

```ts
export interface UIOptions {
  empty?: string;
  className?: string | string[];
  validator?: (
    value,
    formData?,
  ) => Promise<string | true | void> | true | string;
}
```
### InputOptions

```ts
export interface InputOptions extends UIOptions {
  placeholder?: string;
  type?:
    | 'color'
    | 'date'
    | 'datetime-local'
    | 'email'
    | 'month'
    | 'number'
    | 'password'
    | 'range'
    | 'search'
    | 'tel'
    | 'text'
    | 'time'
    | 'url'
    | 'week';
}
```
### SelectOptions

```ts
export interface SelectOptions extends UIOptions {}
```
### UploadOptions

```ts
export interface UploadOptions extends UIOptions {
  acceptType?: string;
  imageType?: 'post' | 'avatar' | 'branding';
}
```
### SwitchOptions

```ts
export interface SwitchOptions extends UIOptions {}
```
### TimezoneOptions

```ts
export interface TimezoneOptions extends UIOptions {
  placeholder?: string;
}
```
### CheckboxOptions

```ts
export interface CheckboxOptions extends UIOptions {}
```
### RadioOptions

```ts
export interface RadioOptions extends UIOptions {}
```
### TextareaOptions

```ts
export interface TextareaOptions extends UIOptions {
  placeholder?: string;
  rows?: number;
}
```
### UIWidget

```ts
export type UIWidget =
  | 'textarea'
  | 'input'
  | 'checkbox'
  | 'radio'
  | 'select'
  | 'upload'
  | 'timezone'
  | 'switch';
```

### UISchema

```ts
export interface UISchema {
  [key: string]: {
    'ui:widget'?: UIWidget;
    'ui:options'?:
      | InputOptions
      | SelectOptions
      | UploadOptions
      | SwitchOptions
      | TimezoneOptions
      | CheckboxOptions
      | RadioOptions
      | TextareaOptions;
  };
}
```

### FormData
```ts
export interface FormValue<T = any> {
  value: T;
  isInvalid: boolean;
  errorMsg: string;
  [prop: string]: any;
}

export interface FormDataType {
  [prop: string]: FormValue;
}
```

## reference

- [json schema](https://json-schema.org/understanding-json-schema/index.html)
- [react-jsonschema-form](https://github.com/rjsf-team/react-jsonschema-form)
- [vue-json-schema-form](https://github.com/lljj-x/vue-json-schema-form/)
