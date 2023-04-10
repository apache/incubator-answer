import {
  ForwardRefRenderFunction,
  forwardRef,
  useImperativeHandle,
  useEffect,
} from 'react';
import { Form, Button, Stack } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import BrandUpload from '../BrandUpload';
import TimeZonePicker from '../TimeZonePicker';
import type * as Type from '@/common/interface';

export interface JSONSchema {
  title: string;
  description?: string;
  required?: string[];
  properties: {
    [key: string]: {
      type: 'string' | 'boolean' | 'number';
      title: string;
      description?: string;
      enum?: Array<string | boolean | number>;
      enumNames?: string[];
      default?: string | boolean | number;
    };
  };
}

export interface BaseUIOptions {
  empty?: string;
  className?: string | string[];
  validator?: (
    value,
    formData?,
  ) => Promise<string | true | void> | true | string;
}
export interface InputOptions extends BaseUIOptions {
  placeholder?: string;
  inputType?:
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
export interface SelectOptions extends BaseUIOptions {}
export interface UploadOptions extends BaseUIOptions {
  acceptType?: string;
  imageType?: Type.UploadType;
}

export interface SwitchOptions extends BaseUIOptions {
  label?: string;
}

export interface TimezoneOptions extends BaseUIOptions {
  placeholder?: string;
}

export interface CheckboxOptions extends BaseUIOptions {}

export interface RadioOptions extends BaseUIOptions {}

export interface TextareaOptions extends BaseUIOptions {
  placeholder?: string;
  rows?: number;
}

export type UIOptions =
  | InputOptions
  | SelectOptions
  | UploadOptions
  | SwitchOptions
  | TimezoneOptions
  | CheckboxOptions
  | RadioOptions
  | TextareaOptions;

export type UIWidget =
  | 'textarea'
  | 'input'
  | 'checkbox'
  | 'radio'
  | 'select'
  | 'upload'
  | 'timezone'
  | 'switch';
export interface UISchema {
  [key: string]: {
    'ui:widget'?: UIWidget;
    'ui:options'?: UIOptions;
  };
}

interface IProps {
  schema: JSONSchema;
  uiSchema?: UISchema;
  formData?: Type.FormDataType;
  hiddenSubmit?: boolean;
  onChange?: (data: Type.FormDataType) => void;
  onSubmit?: (e: React.FormEvent) => void;
}

interface IRef {
  validator: () => Promise<boolean>;
}

/**
 * json schema form
 * @param schema json schema
 * @param uiSchema ui schema
 * @param formData form data
 * @param onChange change event
 * @param onSubmit submit event
 */
const SchemaForm: ForwardRefRenderFunction<IRef, IProps> = (
  {
    schema,
    uiSchema = {},
    formData = {},
    onChange,
    onSubmit,
    hiddenSubmit = false,
  },
  ref,
) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'form',
  });

  const { required = [], properties } = schema;

  // check required field
  const excludes = required.filter((key) => !properties[key]);

  if (excludes.length > 0) {
    console.error(t('not_found_props', { key: excludes.join(', ') }));
  }

  const keys = Object.keys(properties);
  /**
   * Prevent components such as `select` from having default values,
   * which are not generated on `formData`
   */
  const setDefaultValueAsDomBehaviour = () => {
    keys.forEach((k) => {
      const fieldVal = formData[k]?.value;
      const metaProp = properties[k];
      const uiCtrl = uiSchema[k]?.['ui:widget'];
      if (!metaProp || !uiCtrl || fieldVal !== undefined) {
        return;
      }
      if (uiCtrl === 'select' && metaProp.enum?.[0] !== undefined) {
        formData[k] = {
          errorMsg: '',
          isInvalid: false,
          value: metaProp.enum?.[0],
        };
      }
    });
  };
  useEffect(() => {
    setDefaultValueAsDomBehaviour();
  }, [formData]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    const data = {
      ...formData,
      [name]: { ...formData[name], value, isInvalid: false },
    };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value } = e.target;
    const data = {
      ...formData,
      [name]: { ...formData[name], value, isInvalid: false },
    };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  const handleSwitchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = e.target;
    const data = {
      ...formData,
      [name]: { ...formData[name], value: checked, isInvalid: false },
    };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  const requiredValidator = () => {
    const errors: string[] = [];
    required.forEach((key) => {
      if (!formData[key] || !formData[key].value) {
        errors.push(key);
      }
    });
    return errors;
  };

  const syncValidator = () => {
    const errors: Array<{ key: string; msg: string }> = [];
    const promises: Array<{
      key: string;
      promise;
    }> = [];
    keys.forEach((key) => {
      const { validator } = uiSchema[key]?.['ui:options'] || {};
      if (validator instanceof Function) {
        const value = formData[key]?.value;
        promises.push({
          key,
          promise: validator(value, formData),
        });
      }
    });
    return Promise.allSettled(promises.map((item) => item.promise)).then(
      (results) => {
        results.forEach((result, index) => {
          const { key } = promises[index];
          if (result.status === 'rejected') {
            errors.push({
              key,
              msg: result.reason.message,
            });
          }

          if (result.status === 'fulfilled') {
            const msg = result.value;
            if (typeof msg === 'string') {
              errors.push({
                key,
                msg,
              });
            }
          }
        });
        return errors;
      },
    );
  };

  const validator = async (): Promise<boolean> => {
    const errors = requiredValidator();
    if (errors.length > 0) {
      formData = errors.reduce((acc, cur) => {
        acc[cur] = {
          ...formData[cur],
          isInvalid: true,
          errorMsg:
            uiSchema[cur]?.['ui:options']?.empty ||
            `${schema.properties[cur]?.title} ${t('empty')}`,
        };
        return acc;
      }, formData);
      if (onChange instanceof Function) {
        onChange({ ...formData });
      }
      return false;
    }
    const syncErrors = await syncValidator();
    if (syncErrors.length > 0) {
      formData = syncErrors.reduce((acc, cur) => {
        acc[cur.key] = {
          ...formData[cur.key],
          isInvalid: true,
          errorMsg:
            cur.msg || `${schema.properties[cur.key].title} ${t('invalid')}`,
        };
        return acc;
      }, formData);
      if (onChange instanceof Function) {
        onChange({ ...formData });
      }
      return false;
    }
    return true;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const isValid = await validator();
    if (!isValid) {
      return;
    }

    Object.keys(formData).forEach((key) => {
      formData[key].isInvalid = false;
      formData[key].errorMsg = '';
    });
    if (onChange instanceof Function) {
      onChange(formData);
    }
    if (onSubmit instanceof Function) {
      onSubmit(e);
    }
  };

  const handleUploadChange = (name: string, value: string) => {
    const data = { ...formData, [name]: { ...formData[name], value } };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  const handleInputCheck = (
    e: React.ChangeEvent<HTMLInputElement>,
    index: number,
  ) => {
    const { name, checked } = e.currentTarget;
    const freshVal = checked ? schema.properties[name]?.enum?.[index] : '';
    const data = {
      ...formData,
      [name]: {
        ...formData[name],
        value: freshVal,
        isInvalid: false,
      },
    };
    if (onChange instanceof Function) {
      onChange(data);
    }
  };

  useImperativeHandle(ref, () => ({
    validator,
  }));
  return (
    <Form noValidate onSubmit={handleSubmit}>
      {keys.map((key) => {
        const { title, description } = properties[key];
        const { 'ui:widget': widget = 'input', 'ui:options': uiOpt } =
          uiSchema[key] || {};
        if (widget === 'select') {
          return (
            <Form.Group
              key={title}
              controlId={key}
              className={classnames('mb-3', formData[key].hidden && 'd-none')}>
              <Form.Label>{title}</Form.Label>
              <Form.Select
                aria-label={description}
                name={key}
                value={formData[key]?.value || ''}
                onChange={handleSelectChange}
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
              {description && (
                <Form.Text className="text-muted">{description}</Form.Text>
              )}
            </Form.Group>
          );
        }

        if (widget === 'checkbox' || widget === 'radio') {
          return (
            <Form.Group
              key={title}
              className={classnames('mb-3', formData[key].hidden && 'd-none')}
              controlId={key}>
              <Form.Label>{title}</Form.Label>
              <Stack direction="horizontal">
                {properties[key].enum?.map((item, index) => {
                  return (
                    <Form.Check
                      key={String(item)}
                      inline
                      required
                      type={widget}
                      name={key}
                      id={`form-${String(item)}`}
                      label={properties[key].enumNames?.[index]}
                      checked={(formData[key]?.value || '') === item}
                      feedback={formData[key]?.errorMsg}
                      feedbackType="invalid"
                      isInvalid={formData[key].isInvalid}
                      onChange={(e) => handleInputCheck(e, index)}
                    />
                  );
                })}
              </Stack>
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>
              {description && (
                <Form.Text className="text-muted">{description}</Form.Text>
              )}
            </Form.Group>
          );
        }

        if (widget === 'switch') {
          return (
            <Form.Group
              key={title}
              className={classnames('mb-3', formData[key].hidden && 'd-none')}
              controlId={key}>
              <Form.Label>{title}</Form.Label>
              <Form.Check
                required
                id={`switch-${title}`}
                name={key}
                type="switch"
                label={(uiOpt as SwitchOptions)?.label}
                checked={formData[key]?.value || ''}
                feedback={formData[key]?.errorMsg}
                feedbackType="invalid"
                isInvalid={formData[key].isInvalid}
                onChange={handleSwitchChange}
              />
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>
              {description && (
                <Form.Text className="text-muted">{description}</Form.Text>
              )}
            </Form.Group>
          );
        }
        if (widget === 'timezone') {
          return (
            <Form.Group
              key={title}
              className={classnames('mb-3', formData[key].hidden && 'd-none')}
              controlId={key}>
              <Form.Label>{title}</Form.Label>
              <TimeZonePicker
                value={formData[key]?.value || ''}
                name={key}
                onChange={handleSelectChange}
              />
              <Form.Control
                name={key}
                className="d-none"
                isInvalid={formData[key].isInvalid}
              />
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>
              {description && (
                <Form.Text className="text-muted">{description}</Form.Text>
              )}
            </Form.Group>
          );
        }

        if (widget === 'upload') {
          const options: UploadOptions = uiSchema[key]?.['ui:options'] || {};
          return (
            <Form.Group
              key={title}
              className={classnames('mb-3', formData[key].hidden && 'd-none')}
              controlId={key}>
              <Form.Label>{title}</Form.Label>
              <BrandUpload
                type={options.imageType || 'avatar'}
                acceptType={options.acceptType || ''}
                value={formData[key]?.value}
                onChange={(value) => handleUploadChange(key, value)}
              />
              <Form.Control
                name={key}
                className="d-none"
                isInvalid={formData[key].isInvalid}
              />
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>
              {description && (
                <Form.Text className="text-muted">{description}</Form.Text>
              )}
            </Form.Group>
          );
        }

        if (widget === 'textarea') {
          const options: TextareaOptions = uiSchema[key]?.['ui:options'] || {};

          return (
            <Form.Group
              controlId={`form-${key}`}
              key={key}
              className={classnames('mb-3', formData[key].hidden && 'd-none')}>
              <Form.Label>{title}</Form.Label>
              <Form.Control
                as="textarea"
                name={key}
                placeholder={options?.placeholder || ''}
                value={formData[key]?.value || ''}
                onChange={handleInputChange}
                isInvalid={formData[key].isInvalid}
                rows={options?.rows || 3}
                className={classnames(options.className)}
              />
              <Form.Control.Feedback type="invalid">
                {formData[key]?.errorMsg}
              </Form.Control.Feedback>

              {description && (
                <Form.Text className="text-muted">{description}</Form.Text>
              )}
            </Form.Group>
          );
        }

        const options: InputOptions = uiSchema[key]?.['ui:options'] || {};

        return (
          <Form.Group
            controlId={key}
            key={key}
            className={classnames('mb-3', formData[key].hidden && 'd-none')}>
            <Form.Label>{title}</Form.Label>
            <Form.Control
              name={key}
              placeholder={options?.placeholder || ''}
              type={options?.inputType || 'text'}
              value={formData[key]?.value || ''}
              onChange={handleInputChange}
              style={options?.inputType === 'color' ? { width: '6rem' } : {}}
              isInvalid={formData[key].isInvalid}
            />
            <Form.Control.Feedback type="invalid">
              {formData[key]?.errorMsg}
            </Form.Control.Feedback>

            {description && (
              <Form.Text className="text-muted">{description}</Form.Text>
            )}
          </Form.Group>
        );
      })}
      {!hiddenSubmit && (
        <Button variant="primary" type="submit">
          {t('btn_submit')}
        </Button>
      )}
    </Form>
  );
};
export const initFormData = (schema: JSONSchema): Type.FormDataType => {
  const formData: Type.FormDataType = {};
  Object.keys(schema.properties).forEach((key) => {
    const prop = schema.properties[key];
    const defaultVal = prop?.default;
    formData[key] = {
      value: defaultVal,
      isInvalid: false,
      errorMsg: '',
    };
  });
  return formData;
};

export default forwardRef(SchemaForm);
