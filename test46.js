
const error = new Error("エラーメッセージ");
try {
  console.log("正常の動作"　);
  const error = new Error("エラーメッセージ");
} catch(error) {
  console.log("エラー内容："　+ error);
}


const add = (a,b) => {
  const add = a + b;
  if (100 > add) {
    throw new Error("100以上なのでエラー");
  } else {
    console.log("100より下");
  }
};


add(10,101);

window.onerror = function(msg, url, line, col, error) {
    console.log("msg " + msg); // エラーの内容
    console.log("url " + url); // エラーの内容
    console.log("line:col " + line + ":" + col ); // エラーの内容

};
