export interface FormValue<T = any> {
  value: T;
  isInvalid: boolean;
  errorMsg: string;
}

export interface FormDataType {
  [prop: string]: FormValue;
}
