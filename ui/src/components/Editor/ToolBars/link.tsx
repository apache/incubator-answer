import { FC, useEffect, useRef, useState, memo } from 'react';
import { Button, Form, Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Link: FC<IEditorContext> = ({ editor }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'link',
    keyMap: ['Ctrl-L'],
    tip: `${t('link.text')} (Ctrl+L)`,
  };
  const [visible, setVisible] = useState(false);
  const [link, setLink] = useState({
    value: 'http://',
    isInvalid: false,
    errorMsg: '',
  });
  const [name, setName] = useState({
    value: '',
    isInvalid: false,
    errorMsg: '',
  });
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (visible && inputRef.current) {
      inputRef.current.setSelectionRange(0, inputRef.current.value.length);
      inputRef.current.focus();
    }
  }, [visible]);
  const addLink = () => {
    if (!editor) {
      return;
    }
    const text = editor.getSelection();

    setName({ ...name, value: text });

    setVisible(true);
  };
  const handleClick = () => {
    if (!editor) {
      return;
    }
    if (!link.value) {
      setLink({ ...link, isInvalid: true });
      return;
    }
    const newStr = name.value
      ? `[${name.value}](${link.value})`
      : `<${link.value}>`;

    editor.replaceSelection(newStr);

    setVisible(false);

    editor.focus();
    setLink({ ...link, value: '' });
    setName({ ...name, value: '' });
  };
  const onHide = () => setVisible(false);
  const onExited = () => editor?.focus();

  return (
    <ToolItem {...item} click={addLink}>
      <Modal
        show={visible}
        onHide={onHide}
        onExited={onExited}
        fullscreen="sm-down">
        <Modal.Header closeButton>
          <h5 className="mb-0">{t('link.add_link')}</h5>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleClick}>
            <Form.Group controlId="editor.internetSite" className="mb-3">
              <Form.Label>{t('link.form.fields.url.label')}</Form.Label>
              <Form.Control
                ref={inputRef}
                type="text"
                value={link.value}
                onChange={(e) => setLink({ ...link, value: e.target.value })}
                isInvalid={link.isInvalid}
              />
            </Form.Group>

            <Form.Group controlId="editor.internetSiteName" className="mb-3">
              <Form.Label>{t('link.form.fields.name.label')}</Form.Label>
              <Form.Control
                type="text"
                value={name.value}
                onChange={(e) => setName({ ...name, value: e.target.value })}
                isInvalid={name.isInvalid}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="link" onClick={() => setVisible(false)}>
            {t('link.btn_cancel')}
          </Button>
          <Button variant="primary" onClick={handleClick}>
            {t('link.btn_confirm')}
          </Button>
        </Modal.Footer>
      </Modal>
    </ToolItem>
  );
};

export default memo(Link);
