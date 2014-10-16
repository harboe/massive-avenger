(function(win, doc, undefined) {

	ko.validation.rules.pattern.message = 'Invalid.';

	ko.validation.configure({
	    registerExtenders: true,
	    messagesOnModified: true,
	    insertMessages: true,
	    parseInputAttributes: true,
	    messageTemplate: null
	});

	var AccountViewModel = function() {

		this.name = 'User Management'

		this.fullname = ko.observable().extend({ 
			required: true,
			minLength: 6,
		});
		this.email = ko.observable().extend({ 
			email: true 
		});
		this.password = ko.observable().extend({
			required: true,
			minLength: 6,
		});
		this.role = ko.observable().extend({
			required: true,
			min: { value: 1, message: 'This field is required.' },
		});

		this.errors = ko.validation.group(this);
		this.isValid = ko.computed(function() {
			return this.errors().length === 0;
		}, this);

		this.submit = function(viewModel, e) {
			if (viewModel.isValid()) {

				api.account.create(this.fullname(), this.email(), this.password(), this.role(), function(e, err) {

					if (err) {
						console.log(err);
					}

					console.log('hello, world!')

				}, this);

				console.log('yes sir.')
			} else {
            	viewModel.errors.showAllMessages();
			}

		}


	};

	var AdminViewModel = function() {
		this.tabs = ko.observableArray([
			new AccountViewModel()
		]);

		this.selectedTab = ko.observable();


	};

	$().ready(function() {
		console.log('muuhua....')
		ko.applyBindings(new AdminViewModel());
	});	

})(window, document);