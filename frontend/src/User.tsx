import { useState } from "react";
import UserForm from "./components/user/UserForm";
import axios from "axios";
import { showNotification } from "@mantine/notifications";
import UserInfo from "./components/user/UserInfo";
import { Space } from "@mantine/core";

export interface UserIfc {
  id: string;
  login: string;
  created_at: string;
}

interface FormResponseIfc {
  value: string;
  type: "name" | "id";
}

function User() {
  const [user, setUser] = useState<UserIfc | undefined>(undefined);

  function onSubmit(formResponse: FormResponseIfc) {
    const url_base = "https://beta-happening.fdm.com.de/api/user";
    let url =
      formResponse.type === "name" ? url_base + "?name=" : url_base + "?id=";
    url += formResponse.value;

    axios
      .get(url)
      .then((res) => {
        let newUser: UserIfc = {
          id: res.data.id,
          login: res.data.login,
          created_at: res.data.created_at,
        };
        setUser(newUser);
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 404) {
            let message =
              formResponse.type === "name"
                ? `Name '${formResponse.value} not found'`
                : `ID '${formResponse.value}' not found`;
            showNotification({
              message: message,
              color: "red",
            });
          }
        }
      });
  }

  // TODO:
  // - add option to input id

  return (
    <div>
      <UserForm parentSubmit={onSubmit} />
      <Space h="xl" />
      {user !== undefined && <UserInfo user={user}></UserInfo>}
    </div>
  );
}

export default User;
