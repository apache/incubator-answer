import DefaultModal from './Modal';
import confirm, { Config } from './Confirm';
import PicAuthCodeModal from './PicAuthCodeModal';

type ModalType = typeof DefaultModal & {
  confirm: (config: Config) => void;
};
const Modal = DefaultModal as ModalType;

Modal.confirm = function (props: Config) {
  return confirm(props);
};

export default Modal;

export { PicAuthCodeModal };
