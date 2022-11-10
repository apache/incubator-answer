# SchemaForm User Guide

## Introduction

SchemaForm is a component that can be used to render a form based on a [JSON schema](https://json-schema.org/understanding-json-schema/index.html).

## Usage

### Basic Usage

```jsx
import React from 'react';
import { SchemaForm, initFormData, JSONSchema, UISchema } from '@/components';

const schema: JSONSchema = {
  type: 'object',
  properties: {
    name: {
      type: 'string',
      title: 'Name',
    },
    age: {
      type: 'number',
      title: 'Age',
    },
    sex:{
      type: 'boolean',
      title: 'sex',
      enum: [1, 2]
      enumNames: ['male', 'female'],
    }
  },
};

const uiSchema: UISchema = {
  name: {
    'ui:widget': 'input',
  },
  age: {
    'ui:widget': 'input',
    'ui:options': {
      type: 'number'
    }
  },
  sex: {
    'ui:widget': 'radio',
  }
};

// form component

const Form = () => {
  const [formData, setFormData] = useState(initFormData(schema));
  return (
    <SchemaForm
      schema={schema}
      uiSchema={uiSchema}
      formData={formData}
      onChange={console.log}
    />
  );
};
```

## Props

| Property | Description                              | Type                                      | Default |
| -------- | ---------------------------------------- | ----------------------------------------- | ------- |
| schema   | JSON schema                              | [JSONSchema](index.tsx#L9)                | -       |
| uiSchema | UI schema                                | [UISchema](index.tsx#L24)                 | -       |
| formData | Form data                                | [FormData](index.tsx#L66)                 | -       |
| onChange | Callback function when form data changes | (data: [FormData](index.tsx#L66)) => void | -       |
| onSubmit | Callback function when form is submitted | (data: React.FormEvent) => void           | -       |

## reference

- [json schema](https://json-schema.org/understanding-json-schema/index.html)
- [react-jsonschema-form](http://rjsf-team.github.io/react-jsonschema-form/)
- [vue-json-schema-form](https://github.com/lljj-x/vue-json-schema-form/)
