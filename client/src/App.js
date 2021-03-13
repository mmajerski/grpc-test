import "./App.css";

import authClient from "./js/auth_grpc_web_pb";
import authObjs from "./js/auth_pb";

function App() {
  const handleOnClick = (e) => {
    e.preventDefault();

    const client = new authClient.AuthClient("http://localhost:8080");
    const request = new authObjs.LoginRequest();

    request.setEmail("mail@mail.com");
    request.setPassword("test12345");

    client.login(request, { "header-test": "value1" }, (err, resp) => {
      console.log(resp.getToken());
    });
  };

  return (
    <div className="App">
      <button onClick={handleOnClick}>Click</button>
    </div>
  );
}

export default App;
