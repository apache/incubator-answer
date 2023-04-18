export interface UIAction {
  url: string;
  method?: 'get' | 'post' | 'put' | 'delete';
  event?: 'click' | 'change';
  handler?: ({evt, formData, request}) => Promise<void>
}
