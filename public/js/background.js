$.fn.background = function() {
	return this.each(function(index, element) {
		$(element).css('background', 'linear-gradient(rgba(0,0,0,0.4), rgba(0,0,0,0.4)), url(\'' + $(element).data('background') + '\') center / cover');
	});
}
