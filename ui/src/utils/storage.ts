const Storage = {
  get: (key: string): any => {
    const value = localStorage.getItem(key);
    if (value) {
      try {
        const v = JSON.parse(value);
        return v;
      } catch {
        return value;
      }
    }
    return false;
  },
  set: (key: string, value: any): void => {
    if (typeof value === 'string') {
      localStorage.setItem(key, value);
      return;
    }
    localStorage.setItem(key, JSON.stringify(value));
  },
  remove: (key: string): void => {
    localStorage.removeItem(key);
  },
  clear: (): void => {
    localStorage.clear();
  },
};

export default Storage;
