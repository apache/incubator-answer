export interface FormValue<T = any> {
  value: T;
  isInvalid: boolean;
  errorMsg: string;
}

export interface FormDataType {
  [prop: string]: FormValue;
}

export interface Paging {
  page: number;
  page_size?: number;
}

export type ReportType = 'question' | 'answer' | 'comment' | 'user';
export type ReportAction = 'close' | 'flag' | 'review';
export interface ReportParams {
  type: ReportType;
  action: ReportAction;
}

export interface TagBase {
  display_name: string;
  slug_name: string;
}

export interface Tag extends TagBase {
  main_tag_slug_name?: string;
  original_text?: string;
  parsed_text?: string;
}

export interface SynonymsTag extends Tag {
  tag_id: string;
  tag?: string;
}

export interface TagInfo extends TagBase {
  tag_id: string;
  original_text: string;
  parsed_text: string;
  follow_count: number;
  question_count: number;
  is_follower: boolean;
  member_actions;
  created_at?;
  updated_at?;
  main_tag_slug_name?: string;
  excerpt?;
}
export interface QuestionParams {
  title: string;
  content: string;
  html: string;
  tags: Tag[];
}

export interface ListResult<T = any> {
  count: number;
  list: T[];
}

export interface AnswerParams {
  content: string;
  html: string;
  question_id: string;
  id: string;
}

export interface LoginReqParams {
  e_mail: string;
  /** password */
  pass: string;
  captcha_id?: string;
  captcha_code?: string;
}

export interface RegisterReqParams extends LoginReqParams {
  name: string;
}

export interface ModifyPasswordReq {
  old_pass: string;
  pass: string;
}

/** User  */
export interface ModifyUserReq {
  display_name: string;
  username?: string;
  avatar: string;
  bio: string;
  bio_html?: string;
  location: string;
  website: string;
}

export interface UserInfoBase {
  avatar: string;
  username: string;
  display_name: string;
  rank: number;
  website: string;
  location: string;
  ip_info?: string;
  /** 'forbidden' | 'normal' | 'delete'
   */
  status?: string;
  /** roles */
  is_admin?: true;
}

export interface UserInfoRes extends UserInfoBase {
  bio: string;
  bio_html: string;
  create_time?: string;
  /** value = 1 active; value = 2 inactivated
   */
  mail_status: number;
  e_mail?: string;
  [prop: string]: any;
}

export interface AvatarUploadReq {
  file: FormData;
}

export interface ImgCodeReq {
  captcha_id?: string;
  captcha_code?: string;
}

export interface ImgCodeRes {
  captcha_id: string;
  captcha_img: string;
  verify: boolean;
}

export interface PasswordResetReq extends ImgCodeReq {
  e_mail: string;
}

export interface CheckImgReq {
  action: 'login' | 'e_mail' | 'find_pass';
}

export interface SetNoticeReq {
  notice_switch: boolean;
}

export interface QuestionDetailRes {
  id: string;
  title: string;
  content: string;
  html: string;
  tags: any[];
  view_count: number;
  unique_view_count?: number;
  answer_count: number;
  favorites_count: number;
  follow_counts: 0;
  accepted_answer_id: string;
  last_answer_id: string;
  create_time: string;
  update_time: string;
  user_info: UserInfoBase;
  answered: boolean;
  collected: boolean;

  [prop: string]: any;
}

export interface AnswersReq extends Paging {
  order?: 'default' | 'updated';
  question_id: string;
}

export interface AnswerItem {
  id: string;
  question_id: string;
  content: string;
  html: string;
  create_time: string;
  update_time: string;
  user_info: UserInfoBase;

  [prop: string]: any;
}

export interface PostAnswerReq {
  content: string;
  html: string;
  question_id: string;
}

export interface PageUser {
  id;
  displayName;
  userName?;
  avatar_url?;
}

export interface LangsType {
  label: string;
  value: string;
}

/**
 * @description interface for Question
 */
export type QuestionOrderBy =
  | 'newest'
  | 'active'
  | 'frequent'
  | 'score'
  | 'unanswered';

export interface QueryQuestionsReq extends Paging {
  order: QuestionOrderBy;
  tags?: string[];
}

export type AdminQuestionStatus = 'available' | 'closed' | 'deleted';

export type AdminContentsFilterBy = 'normal' | 'closed' | 'deleted';

export interface AdminContentsReq extends Paging {
  status: AdminContentsFilterBy;
}

/**
 * @description interface for Answer
 */
export type AdminAnswerStatus = 'available' | 'deleted';

/**
 * @description interface for Users
 */
export type UserFilterBy = 'all' | 'inactive' | 'suspended' | 'deleted';

/**
 * @description interface for Flags
 */
export type FlagStatus = 'pending' | 'completed';
export type FlagType = 'all' | 'question' | 'answer' | 'comment';
export interface AdminFlagsReq extends Paging {
  status: FlagStatus;
  object_type: FlagType;
}

/**
 * @description interface for Admin Settings
 */
export interface AdminSettingsGeneral {
  name: string;
  short_description: string;
  description: string;
}

export interface AdminSettingsInterface {
  logo: string;
  language: string;
  theme: string;
}

export interface SiteSettings {
  general: AdminSettingsGeneral;
  interface: AdminSettingsInterface;
}
/**
 * @description interface for Activity
 */
export interface FollowParams {
  is_cancel: boolean;
  object_id: string;
}

/**
 * @description search request params
 */
export interface SearchParams {
  q: string;
  order: string;
  page: number;
  size?: number;
}

/**
 * @description search response data
 */
export interface SearchResItem {
  object_type: string;
  object: {
    id: string;
    title: string;
    excerpt: string;
    created_at: number;
    user_info: UserInfoBase;
    vote_count: number;
    answer_count: number;
    accepted: boolean;
    tags: TagBase[];
  };
}
export interface SearchRes extends ListResult<SearchResItem> {
  extra: any;
}
