import DefaultModal from './Modal';
import confirm, { Config } from './Confirm';
import PicAuthCodeModal from './PicAuthCodeModal';
import LoginToContinueModal from './LoginToContinueModal';

type ModalType = typeof DefaultModal & {
  confirm: (config: Config) => void;
};
const Modal = DefaultModal as ModalType;

Modal.confirm = function (props: Config) {
  return confirm(props);
};

export default Modal;

export { PicAuthCodeModal };
export { LoginToContinueModal };
