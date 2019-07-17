namespace StrongForce.Core.Tests.Mocks
{
	public class SetFavoriteNumberAction : Action
	{
		public SetFavoriteNumberAction(Address target, int number)
			: base(target)
		{
			this.Number = number;
		}

		public int Number { get; }
	}
}