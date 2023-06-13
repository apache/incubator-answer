import React, {
  ForwardRefRenderFunction,
  forwardRef,
  useImperativeHandle,
  useEffect,
} from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { isEmpty } from 'lodash';
import classnames from 'classnames';

import type * as Type from '@/common/interface';

import type {
  JSONSchema,
  FormProps,
  FormRef,
  BaseUIOptions,
  FormKit,
} from './types';
import {
  Legend,
  Select,
  Check,
  Switch,
  Timezone,
  Upload,
  Textarea,
  Input,
  Button as SfButton,
} from './components';

export * from './types';

/**
 * TODO:
 *  - [!] Standardised `Admin/Plugins/Config/index.tsx` method for generating dynamic form configurations.
 *  - Normalize and document `formData[key].hidden && 'd-none'`
 *  - Normalize and document `hiddenSubmit`
 *  - Improving field hints for `formData`
 *  - Optimise form data updates
 *    * Automatic field type conversion
 *    * Dynamic field generation
 */

/**
 * json schema form
 * @param schema json schema
 * @param uiSchema ui schema
 * @param formData form data
 * @param onChange change event
 * @param onSubmit submit event
 */
const SchemaForm: ForwardRefRenderFunction<FormRef, FormProps> = (
  {
    schema,
    uiSchema = {},
    refreshConfig,
    formData,
    onChange,
    onSubmit,
    hiddenSubmit = false,
  },
  ref,
) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'form',
  });
  const { required = [], properties = {} } = schema || {};
  // check required field
  const excludes = required.filter((key) => !properties[key]);
  if (excludes.length > 0) {
    console.error(t('not_found_props', { key: excludes.join(', ') }));
  }
  formData ||= {};
  const keys = Object.keys(properties);
  /**
   * Prevent components such as `select` from having default values,
   * which are not generated on `formData`
   */
  const setDefaultValueAsDomBehaviour = () => {
    keys.forEach((k) => {
      const fieldVal = formData![k]?.value;
      const metaProp = properties[k];
      const uiCtrl = uiSchema[k]?.['ui:widget'];
      if (!metaProp || !uiCtrl || fieldVal !== undefined) {
        return;
      }
      if (uiCtrl === 'select' && metaProp.enum?.[0] !== undefined) {
        formData![k] = {
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

  const formKitWithContext: FormKit = {
    refreshConfig() {
      if (typeof refreshConfig === 'function') {
        refreshConfig();
      }
    },
  };

  /**
   * Form validation
   * - Currently only dynamic forms are in use, the business form validation has been handed over to the server
   */
  const requiredValidator = () => {
    const errors: string[] = [];
    required.forEach((key) => {
      if (!formData![key] || !formData![key].value) {
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
        const value = formData![key]?.value;
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
          ...formData![cur],
          isInvalid: true,
          errorMsg:
            uiSchema[cur]?.['ui:options']?.empty ||
            `${properties[cur]?.title} ${t('empty')}`,
        };
        return acc;
      }, formData || {});
      if (onChange instanceof Function) {
        onChange({ ...formData });
      }
      return false;
    }
    const syncErrors = await syncValidator();
    if (syncErrors.length > 0) {
      formData = syncErrors.reduce((acc, cur) => {
        acc[cur.key] = {
          ...formData![cur.key],
          isInvalid: true,
          errorMsg: cur.msg || `${properties[cur.key].title} ${t('invalid')}`,
        };
        return acc;
      }, formData || {});
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

    Object.keys(formData!).forEach((key) => {
      formData![key].isInvalid = false;
      formData![key].errorMsg = '';
    });
    if (onChange instanceof Function) {
      onChange(formData!);
    }
    if (onSubmit instanceof Function) {
      onSubmit(e);
    }
  };

  useImperativeHandle(ref, () => ({
    validator,
  }));
  if (!formData || !schema || isEmpty(schema.properties)) {
    return null;
  }

  return (
    <Form noValidate onSubmit={handleSubmit}>
      {keys.map((key) => {
        const {
          title,
          description,
          enum: enumValues = [],
          enumNames = [],
        } = properties[key];
        const { 'ui:widget': widget = 'input', 'ui:options': uiOpt } =
          uiSchema?.[key] || {};
        formData ||= {};
        const fieldState = formData[key];
        const uiSimplify = widget === 'legend' || uiOpt?.simplify;
        let groupClassName: BaseUIOptions['fieldClassName'] = uiOpt?.simplify
          ? 'mb-2'
          : 'mb-3';
        if (widget === 'legend') {
          groupClassName = 'mb-0';
        }
        if (uiOpt?.fieldClassName) {
          groupClassName = uiOpt.fieldClassName;
        }
        const readOnly = uiOpt?.readOnly || false;
        return (
          <Form.Group
            key={title}
            controlId={key}
            className={classnames(
              groupClassName,
              formData[key].hidden ? 'd-none' : null,
            )}>
            {/* Uniform processing `label` */}
            {title && !uiSimplify ? <Form.Label>{title}</Form.Label> : null}
            {/* Handling of individual specific controls */}
            {widget === 'legend' ? <Legend title={title} /> : null}
            {widget === 'select' ? (
              <Select
                desc={description}
                fieldName={key}
                onChange={onChange}
                enumValues={enumValues}
                enumNames={enumNames}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'radio' || widget === 'checkbox' ? (
              <Check
                type={widget}
                fieldName={key}
                onChange={onChange}
                enumValues={enumValues}
                enumNames={enumNames}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'switch' ? (
              <Switch
                title={title}
                label={uiOpt && 'label' in uiOpt ? uiOpt.label : ''}
                fieldName={key}
                onChange={onChange}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'timezone' ? (
              <Timezone
                fieldName={key}
                onChange={onChange}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'upload' ? (
              <Upload
                type={
                  uiOpt && 'imageType' in uiOpt ? uiOpt.imageType : undefined
                }
                acceptType={
                  uiOpt && 'acceptType' in uiOpt ? uiOpt.acceptType : ''
                }
                fieldName={key}
                onChange={onChange}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'textarea' ? (
              <Textarea
                placeholder={
                  uiOpt && 'placeholder' in uiOpt ? uiOpt.placeholder : ''
                }
                rows={uiOpt && 'rows' in uiOpt ? uiOpt.rows : 3}
                className={uiOpt && 'className' in uiOpt ? uiOpt.className : ''}
                fieldName={key}
                onChange={onChange}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'input' ? (
              <Input
                type={uiOpt && 'inputType' in uiOpt ? uiOpt.inputType : 'text'}
                placeholder={
                  uiOpt && 'placeholder' in uiOpt ? uiOpt.placeholder : ''
                }
                fieldName={key}
                onChange={onChange}
                formData={formData}
                readOnly={readOnly}
              />
            ) : null}
            {widget === 'button' ? (
              <SfButton
                fieldName={key}
                text={uiOpt && 'text' in uiOpt ? uiOpt.text : ''}
                action={uiOpt && 'action' in uiOpt ? uiOpt.action : undefined}
                formKit={formKitWithContext}
                readOnly={readOnly}
                variant={
                  uiOpt && 'variant' in uiOpt ? uiOpt.variant : undefined
                }
                size={uiOpt && 'size' in uiOpt ? uiOpt.size : undefined}
              />
            ) : null}
            {/* Unified handling of `Feedback` and `Text` */}
            <Form.Control.Feedback type="invalid">
              {fieldState?.errorMsg}
            </Form.Control.Feedback>
            {description ? (
              <Form.Text className="text-muted">{description}</Form.Text>
            ) : null}
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
  const props: JSONSchema['properties'] = schema?.properties || {};
  Object.keys(props).forEach((key) => {
    const prop = props[key];
    const defaultVal = prop?.default;

    formData[key] = {
      value: defaultVal,
      isInvalid: false,
      errorMsg: '',
    };
  });
  return formData;
};

export const mergeFormData = (
  target: Type.FormDataType | null,
  origin: Type.FormDataType | null,
) => {
  if (!target) {
    return origin;
  }
  if (!origin) {
    return target;
  }
  Object.keys(target).forEach((k) => {
    const oi = origin[k];
    if (oi && oi.value !== undefined) {
      target[k] = {
        value: oi.value,
        isInvalid: false,
        errorMsg: '',
      };
    }
  });
  return target;
};

export default forwardRef(SchemaForm);
