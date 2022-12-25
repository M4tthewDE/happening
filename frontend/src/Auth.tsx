function Auth() {
  const state = document.location.hash.split("&")[2].substring(6);
  const token = document.location.hash.substring(14).split("&")[0];

  if (localStorage.getItem("state") !== state) {
    return <>ERROR</>;
  }

  localStorage.setItem("user_token", token);
  window.location.href = `${process.env.REACT_APP_PATH}`;
  return <>SUCCESS</>;
}

export default Auth;
