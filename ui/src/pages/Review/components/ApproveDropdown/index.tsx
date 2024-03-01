import { useState } from 'react';
import { Dropdown, Button } from 'react-bootstrap';

import EditPostModal from '../EditPostModal';

const Index = () => {
  const [showEditPostModal, setShowEditPostModal] = useState(false);

  const handleEditPostModalState = () => {
    setShowEditPostModal(!showEditPostModal);
  };

  return (
    <div>
      <Dropdown>
        <Dropdown.Toggle
          as={Button}
          variant="outline-primary"
          id="dropdown-basic">
          Approve
        </Dropdown.Toggle>

        <Dropdown.Menu>
          <Dropdown.Item href="#/action-1">Deactivate user</Dropdown.Item>
          <Dropdown.Item href="#/action-2">Suspend user</Dropdown.Item>
          <Dropdown.Item href="#/action-3">Delete user</Dropdown.Item>
        </Dropdown.Menu>
      </Dropdown>
      <EditPostModal
        visible={showEditPostModal}
        handleClose={handleEditPostModalState}
      />
    </div>
  );
};

export default Index;
