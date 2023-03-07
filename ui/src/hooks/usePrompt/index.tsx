import { useCallback } from 'react';
import {
  useBeforeUnload,
  unstable_usePrompt as usePrompt,
} from 'react-router-dom';
import { useTranslation } from 'react-i18next';

// https://gist.github.com/chaance/2f3c14ec2351a175024f62fd6ba64aa6
// The link above is an example of implementing usePrompt with useBlocker.
interface PromptProps {
  when: boolean;
  beforeUnload?: boolean;
}

const usePromptWithUnload = ({
  when = false,
  beforeUnload = true,
}: PromptProps) => {
  const { t } = useTranslation('translation', { keyPrefix: 'prompt' });
  usePrompt({
    when,
    message: `${t('leave_page')} ${t('changes_not_save')}`,
  });

  useBeforeUnload(
    useCallback(
      (event) => {
        if (beforeUnload && when) {
          const msg = t('changes_not_save');
          event.preventDefault();
          event.returnValue = msg;
        }
      },
      [when, beforeUnload],
    ),
    { capture: true },
  );
};

export default usePromptWithUnload;
