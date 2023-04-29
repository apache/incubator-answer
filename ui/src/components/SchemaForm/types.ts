/**
 * A few notes on button control：
 *  - Mainly used to send a request and notify the result of the request, and to update the data as required
 *  - A scenario where a message notification is displayed directly after a click without sending a request, implementing a dedicated control
 *  - Scenarios where the page jumps directly after a click without sending a request, implementing a dedicated control
 *
 * @field url : Target address for sending requests
 * @field method : Method for sending requests, default `get`
 * @field callback: Button event handler function that will fully take over the button events when this field is configured
 *                 *** Incomplete, DO NOT USE ***
 * @field loading: Set button loading information
 * @field notify: Configure how button action processing results are prompted
 */

/**
 * TODO:
 *  - Refining the type of `notify.options`
 */
export interface UIAction {
  url: string;
  method?: 'get' | 'post' | 'put' | 'delete';
  callback?: () => Promise<void>;
  loading?: {
    text: string;
    state?:  'none' | 'pending' | 'completed';
  }
  toastMessage?: boolean;
}
