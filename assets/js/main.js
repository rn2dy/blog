$(function(){
  $('#subscribe').click(function(e){
    e.preventDefault();
    $('#subscribe-form').modal({
      fadeDuration: 250
    });
  });
});
