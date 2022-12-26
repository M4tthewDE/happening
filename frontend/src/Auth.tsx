function Auth() {
  const token = document.location.hash.substring(14).split("&")[0];

  localStorage.setItem("user_token", token);
  window.location.href = `${process.env.REACT_APP_PATH}`;
  return <>SUCCESS</>;
}

export default Auth;
