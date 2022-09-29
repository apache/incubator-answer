import React from 'react';

import ModifyEmail from './components/ModifyEmail';
import ModifyPassword from './components/ModifyPass';

const Index = () => {
  return (
    <>
      <ModifyEmail />
      <ModifyPassword />
    </>
  );
};

export default React.memo(Index);
