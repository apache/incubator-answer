import { FC } from 'react';
import { Form, Button, Stack } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import BrandUpload from '../BrandUpload';
import TimeZonePicker from '../TimeZonePicker';
import type * as Type from '@/common/interface';

export interface JSONSchema {
  title: string;
  description?: string;
  required?: string[];
  properties: {
    [key: string]: {
      type: 'string' | 'boolean';
      title: string;
      description?: string;
      enum?: Array<string | boolean>;
      enumNames?: string[];
      default?: string | boolean;
    };
  };
}
export interface UISchema {
  [key: string]: {
    'ui:widget'?:
      | 'textarea'
      | 'text'
      | 'checkbox'
      | 'radio'
      | 'select'
      | 'upload'
      | 'timezone'
      | 'switch';
    'ui:options'?: {
      rows?: number;
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
      empty?: string;
      invalid?: string;
      validator?: (value) => boolean;
      textRender?: () => React.ReactElement;
      imageType?: 'avatar' | 'logo' | 'mobile_logo' | 'square_icon' | 'favicon';
    };
  };
}

interface IProps {
  schema: JSONSchema;
  uiSchema?: UISchema;
  formData?: Type.FormDataType;
  onChange?: (data: Type.FormDataType) => void;
  onSubmit: (e: React.FormEvent) => void;
}

/**
 * json schema form
 * @param schema json schema
 * @param uiSchema ui schema
 * @param formData form data
 * @param onChange change event
 * @param onSubmit submit event
 */
const SchemaForm: FC<IProps> = ({
  schema,
  uiSchema = {},
  formData = {},
  onChange,
  onSubmit,
}) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'form',
  });
  const { properties } = schema;
  const keys = Object.keys(properties);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    const data = { ...formData, [name]: { ...formData[name], value } };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  const requiredValidator = () => {
    const required = schema.required || [];
    const errors: string[] = [];
    required.forEach((key) => {
      if (!formData[key] || !formData[key].value) {
        errors.push(key);
      }
    });
    return errors;
  };

  const syncValidator = () => {
    const errors: string[] = [];
    keys.forEach((key) => {
      const { validator } = uiSchema[key]?.['ui:options'] || {};
      if (validator instanceof Function) {
        const value = formData[key]?.value;
        if (!validator(value)) {
          errors.push(key);
        }
      }
    });
    return errors;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const errors = requiredValidator();
    if (errors.length > 0) {
      formData = errors.reduce((acc, cur) => {
        acc[cur] = {
          ...formData[cur],
          isInvalid: true,
          errorMsg:
            uiSchema[cur]['ui:options']?.empty ||
            `${schema.properties[cur].title} ${t('form.empty')}`,
        };
        return acc;
      }, formData);
      if (onChange instanceof Function) {
        onChange(formData);
      }
      return;
    }
    const syncErrors = syncValidator();
    if (syncErrors.length > 0) {
      formData = syncErrors.reduce((acc, cur) => {
        acc[cur] = {
          ...formData[cur],
          isInvalid: true,
          errorMsg:
            uiSchema[cur]['ui:options']?.invalid ||
            `${schema.properties[cur].title} ${t('form.invalid')}`,
        };
        return acc;
      }, formData);
      if (onChange instanceof Function) {
        onChange(formData);
      }
      return;
    }
    Object.keys(formData).forEach((key) => {
      formData[key].isInvalid = false;
      formData[key].errorMsg = '';
    });
    if (onChange instanceof Function) {
      onChange(formData);
    }
    onSubmit(e);
  };

  const handleUploadChange = (name: string, value: string) => {
    const data = { ...formData, [name]: { ...formData[name], value } };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  return (
    <Form noValidate onSubmit={handleSubmit}>
      {keys.map((key) => {
        const { title, description } = properties[key];
        const { 'ui:widget': widget = 'input', 'ui:options': options = {} } =
          uiSchema[key] || {};
        if (widget === 'select') {
          return (
            <Form.Group key={title} controlId={key} className="mb-3">
              <Form.Label>{title}</Form.Label>
              <Form.Select
                aria-label={description}
                isInvalid={formData[key].isInvalid}>
                {properties[key].enum?.map((item, index) => {
                  return (
                    <option value={String(item)} key={String(item)}>
                      {properties[key].enumNames?.[index]}
                    </option>
                  );
                })}
              </Form.Select>
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>
              <Form.Text className="text-muted">{description}</Form.Text>
            </Form.Group>
          );
        }
        if (widget === 'checkbox' || widget === 'radio') {
          return (
            <Form.Group key={title} className="mb-3" controlId={key}>
              <Form.Label>{title}</Form.Label>
              <Stack direction="horizontal">
                {properties[key].enum?.map((item, index) => {
                  return (
                    <Form.Check
                      key={String(item)}
                      inline
                      required
                      type={widget}
                      name={title}
                      id={String(item)}
                      label={properties[key].enumNames?.[index]}
                      checked={formData[key]?.value === item}
                      feedback={formData[key]?.errorMsg}
                      feedbackType="invalid"
                      isInvalid={formData[key].isInvalid}
                    />
                  );
                })}
              </Stack>
              <Form.Text className="text-muted">{description}</Form.Text>
            </Form.Group>
          );
        }

        if (widget === 'switch') {
          return (
            <Form.Group key={title} className="mb-3" controlId={key}>
              <Form.Label>{title}</Form.Label>
              <Form.Check
                required
                id={title}
                type="switch"
                label={title}
                feedback={formData[key]?.errorMsg}
                feedbackType="invalid"
                isInvalid={formData[key].isInvalid}
              />
              <Form.Text className="text-muted">{description}</Form.Text>
            </Form.Group>
          );
        }
        if (widget === 'timezone') {
          return (
            <Form.Group key={title} className="mb-3" controlId={key}>
              <Form.Label>{title}</Form.Label>
              <TimeZonePicker
                value={formData[key]?.value}
                onChange={handleInputChange}
              />
              <Form.Text className="text-muted">{description}</Form.Text>
            </Form.Group>
          );
        }

        if (widget === 'upload') {
          return (
            <Form.Group key={title} className="mb-3" controlId={key}>
              <Form.Label>{title}</Form.Label>
              <BrandUpload
                type={options.imageType || 'avatar'}
                value={formData[key]?.value}
                onChange={(value) => handleUploadChange(key, value)}
              />
              <Form.Text className="text-muted">{description}</Form.Text>
            </Form.Group>
          );
        }

        if (widget === 'textarea') {
          return (
            <Form.Group controlId={key} key={key} className="mb-3">
              <Form.Label>{title}</Form.Label>
              <Form.Control
                as="textarea"
                name={key}
                placeholder={options?.placeholder || ''}
                type={options?.type || 'text'}
                value={formData[key]?.value}
                onChange={handleInputChange}
                isInvalid={formData[key].isInvalid}
                rows={options?.rows || 3}
              />
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>

              <Form.Text className="text-muted">{description}</Form.Text>
            </Form.Group>
          );
        }
        return (
          <Form.Group controlId={key} key={key} className="mb-3">
            <Form.Label>{title}</Form.Label>
            <Form.Control
              name={key}
              placeholder={options?.placeholder || ''}
              type={options?.type || 'text'}
              value={formData[key]?.value}
              onChange={handleInputChange}
              style={options?.type === 'color' ? { width: '6rem' } : {}}
              isInvalid={formData[key].isInvalid}
            />
            <Form.Control.Feedback type="invalid">
              {formData[key]?.errorMsg}
            </Form.Control.Feedback>

            <Form.Text className="text-muted">{description}</Form.Text>
          </Form.Group>
        );
      })}
      <Button variant="primary" type="submit">
        {t('btn_submit')}
      </Button>
    </Form>
  );
};
export const initFormData = (schema: JSONSchema): Type.FormDataType => {
  const formData: Type.FormDataType = {};
  Object.keys(schema.properties).forEach((key) => {
    formData[key] = {
      value: '',
      isInvalid: false,
      errorMsg: '',
    };
  });
  return formData;
};

export default SchemaForm;
