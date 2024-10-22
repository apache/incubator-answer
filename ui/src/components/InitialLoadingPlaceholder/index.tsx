// Same as spin in `public/index.html`

import './index.scss';

function InitialLoadingPlaceholder() {
  return (
    <div className="InitialLoadingPlaceholder">
      <div className="InitialLoadingPlaceholder-spinnerContainer">
        <div className="InitialLoadingPlaceholder-spinner" />
      </div>
    </div>
  );
}

export default InitialLoadingPlaceholder;
