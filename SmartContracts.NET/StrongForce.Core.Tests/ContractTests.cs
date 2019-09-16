using System;
using System.Collections.Generic;
using StrongForce.Core.Permissions;
using StrongForce.Core.Tests.Mocks;
using Xunit;

namespace StrongForce.Core.Tests
{
	public class ContractTests
	{
		private readonly IAddressFactory addressFactory = new RandomAddressFactory();

		[Fact]
		public void Receive_WhenPassedNull_ThrowsArgumentNullException()
		{
			var (contract, receiver) = BaseContract.Create(typeof(FavoriteNumberContract), this.addressFactory.Create(), new Dictionary<string, object>(), default(ContractHandlers));

			Assert.Throws<ArgumentNullException>(() => receiver.Invoke(null));
		}

		[Fact]
		public void Receive_WhenReceivedSupportedAction_ReturnsTrue()
		{
			var (contract, receiver) = BaseContract.Create(typeof(FavoriteNumberContract), this.addressFactory.Create(), new Dictionary<string, object>() { { "User", null } }, default(ContractHandlers));

			receiver.Invoke(new Message(
				contract.Address,
				contract.Address,
				contract.Address,
				SetFavoriteNumberAction.Type,
				new Dictionary<string, object>()
				{
					{ SetFavoriteNumberAction.Number, 0 },
				}));
		}
	}
}