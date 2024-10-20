import { Spinner } from 'react-bootstrap';

function InitialLoadingPlaceholder() {
  return (
    <div
      style={{
        flexGrow: 1,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}>
      <Spinner />
      <span style={{ marginLeft: 8 }}>Initializing</span>
    </div>
  );
}

export default InitialLoadingPlaceholder;
