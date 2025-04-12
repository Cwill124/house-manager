document.addEventListener('DOMContentLoaded',function(){
  handleLoginForm();
});

function handleLoginForm(){
 const form = document.getElementById("login-form");
  if(form){
    form.addEventListener('submit',function(e){
      e.preventDefault();
      const username = document.getElementById('username').value;
      const password = document.getElementById('password').value;
      console.log(username,password);
      login(username,password);

    });

  }
}

async function login(username,password){
  const data = {
    "username": username,
    "password" : password
  };

  try {
    const response = await fetch('/api/login', {
      method : 'POST',
      headers :{
        'Content-Type' : 'application/json',

      },
      body: JSON.stringify(data),
    });

    const responseData = await response.json();
    localStorage.setItem('jwt',responseData.access_token);
    localStorage.setItem('refresh',responseData.refresh_token);
  } catch(err){
    console.error(err)
  }
}
