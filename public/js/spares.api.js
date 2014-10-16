(function(win, doc, undefined) {

	win.api = function() {

		var ajax = {
			execute: function(url, type, data, callback, context) {

				var apply = function(e, err) {
					if (callback) {
						callback.apply(context || {}, [e, err]);
					}
					if (err)
						console.log(err);
				};
				var errorCallback = function(e) { apply(null, e); };
				var successCallback = function(e) { apply(e, null); };

				var options = {url: url, type: type, error: errorCallback, sucess: successCallback};

				if (data) {
					options.data = data;
				}

				$.ajax(options);
			},
			get: function(url, callback, context) {
				this.execute(url, 'GET', null, callback, context);

			},
			post: function(url, data, callback, context) {
				this.execute(url, 'POST', data, callback, context);
			},
			del: function(url, callback, context) {
				this.execute(url, 'DELETE', null, callback, context);

			},
			put: function(url, data, callback, context) {
				this.execute(url, 'PUT', data, callback, context);				
			}
		}

		var _accountApi = {
			create: function(name, email, password, role, callback, context) {
				var data = {
					fullname: name,
					email: email,
					password: password, 
					role: role
				};

				ajax.post('/api/admin/account', data, callback, context);
			},
			update: function(id, name, email, role, callback, context) {
				var data = {
					fullname: name,
					email: email,
					role: role
				};

				ajax.put('/api/admin/account/' + id, data, callback, context);
			},
			get: function(id, callback, context) {
				ajax.get('/api/admin/account/' + id, callback, context);
			},
			getAll: function(callback, context) {
				ajax.get('/api/admin/account', callback, context);
			}
			


		};

		return {
			account: _accountApi
		};

	}();



})(window, document);