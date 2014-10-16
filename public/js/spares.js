(function(win, doc, undefined) {

	var localStorage = win.localStorage;

	$().ready(function() {
		$('.list-group-item').click(function(e) {

	        // show content.
	        $('#areas > div').hide();
	        $(this.hash).show();
	        // active link.
	        var li = $(this);
	        li.parent().find('.list-group-item').removeClass('active');
	        li.addClass('active');

	        console.log('muuha...')

	        e.preventDefault();
	        return false;
	    });
	});

})(window, document);