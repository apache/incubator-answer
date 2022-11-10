# SchemaForm User Guide

## Introduction

SchemaForm is a component that can be used to render a form based on a JSON schema.

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

| Property | Description                              | Type                         | Default |
| -------- | ---------------------------------------- | ---------------------------- | ------- |
| schema   | JSON schema                              | [JSONSchema]()               | -       |
| uiSchema | UI schema                                | [UISchema]()                 | -       |
| formData | Form data                                | [FormData]()                 | -       |
| onChange | Callback function when form data changes | (data: [FormData]()) => void | -       |
| onSubmit | Callback function when form is submitted | (data: [FormData]()) => void | -       |
