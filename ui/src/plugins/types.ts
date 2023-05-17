import { UIOptions, UIWidget } from '@/components/SchemaForm';

export interface PluginOption {
  label: string;
  value: string;
}

export interface PluginItem {
  name: string;
  type: UIWidget;
  title: string;
  description: string;
  ui_options?: UIOptions;
  options?: PluginOption[];
  value?: string;
  required?: boolean;
}

export interface PluginConfig {
  name: string;
  slug_name: string;
  config_fields: PluginItem[];
}
