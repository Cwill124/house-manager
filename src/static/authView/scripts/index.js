document.addEventListener('DOMContentLoaded',function(){
  handleLoginForm();
  clearLocalStorage();
});

function handleLoginForm(){
 const form = document.getElementById("login-form");
  if(form){
    form.addEventListener('submit',function(e){
      e.preventDefault();
      const username = document.getElementById('username').value;
      const password = document.getElementById('password').value;
      clearLocalStorage();
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
    localStorage.setItem("userId",responseData.userId);
    window.location.href = responseData.redirectTo;
  } catch(err){
    console.error(err)
  }
}

function clearLocalStorage() {
  localStorage.clear()
}
