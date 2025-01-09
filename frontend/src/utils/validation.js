export const isValidURL = (url) => {
    const pattern = new RegExp(
      "^(https?:\\/\\/)?([a-zA-Z0-9.-]+\\.[a-zA-Z]{2,})(:[0-9]{1,5})?(\\/.*)?$",
      "i"
    );
    return pattern.test(url);
  };
  